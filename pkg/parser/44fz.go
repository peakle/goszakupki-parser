package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/peakle/goszakupki-parser/pkg/manager"
	"github.com/peakle/goszakupki-parser/pkg/provider"
	"github.com/peakle/goszakupki-parser/pkg/proxy"
	"github.com/urfave/cli"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// ProcessLot44 collect data about new lots for 44-FZ
func ProcessLot44(_ *cli.Context) error {
	fmt.Println("Start time: ", time.Now().Format("2006-01-02 15:04"))

	var proxyChan = make(chan string, 3000)
	var doneChan = make(chan struct{}, 2)
	defer func() {
		doneChan <- struct{}{} // for proxy
		doneChan <- struct{}{} // for upserts

		close(doneChan)
		close(proxyChan)
	}()

	go proxy.LoadProxy(proxyChan, doneChan)

	lotChan := make(chan *provider.Purchase, 1000)
	upsertWg := &sync.WaitGroup{}

	upsertWg.Add(1)
	go upsertLot(lotChan, doneChan, upsertWg)

	concurCh := make(chan struct{}, 10) // increase for more parallelism
	lot44Worker(0, lotChan, <-proxyChan, concurCh)

	upsertWg.Wait()
	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
}

func lot44Worker(regNumber int, lotCh chan<- *provider.Purchase, proxy string, concurCh chan struct{}) {
	concurCh <- struct{}{}
	defer func() {
		<-concurCh
	}()

	var err error

	client := fasthttp.Client{}
	if proxy != "" {
		client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(proxy, provider.DefaultTimeout)
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(fmt.Sprintf(provider.URIPatternFZ44Purchase, regNumber))
	req.Header.SetUserAgent(provider.UserAgent)
	req.SetConnectionClose()

	err = client.DoTimeout(req, resp, provider.DefaultTimeout)
	if err != nil {
		log.Println("on lot44Logic: on DoTimeout: " + err.Error())
		return
	}

	var dto *provider.Dto44fz
	err = json.Unmarshal(resp.Body(), dto)
	if dto == nil {
		if err != nil {
			log.Println("on lot44Logic: on unmarshal: " + err.Error())
		}

		return
	}

	{
		purchase := &provider.Purchase{
			Id:           dto.Dto.HeaderBlock.PurchaseNumber,
			Fz:           dto.Dto.HeaderBlock.PlacingWayFZ,
			Customer:     dto.Dto.HeaderBlock.OrganizationPublishName,
			CustomerLink: dto.Dto.HeaderBlock.OrganizationPublishLink,
			// CustomerInn: dto.Dto.Inn,
			CustomerRegion: dto.Dto.OrganizationDefinitionSupplierBlock.Location,
			// BiddingRegion: ,
			// CustomerActivityField: dto.Dto.HeaderBlock.PurchaseObjectName,
			BiddingVolume: fmt.Sprintf("%.6f", dto.Dto.InitialContractPriceBlock.InitialContractPrice),
			// BiddingCount: ,
			PurchaseTarget:        dto.Dto.HeaderBlock.PurchaseObjectName,
			RegistryBiddingNumber: dto.Dto.HeaderBlock.PurchaseNumber,
			ContractPrice:         fmt.Sprintf("%.6f", dto.Dto.InitialContractPriceBlock.InitialContractPrice),
			PublishedAt:           time.Unix(dto.Dto.ProcedurePurchaseBlock.StartDateTime, 0).Format("02-01-2006"), // maybe error need treem 3 digits from rigth end
			RequisitionDeadlineAt: time.Unix(dto.Dto.ProcedurePurchaseBlock.EndDateTime, 0).Format("02-01-2006"),   // maybe error need treem 3 digits from rigth end
			ContractStartAt:       "",                                                                              //TODO
			ContractEndAt:         "",                                                                              //TODO
			Playground:            dto.Dto.GeneralInformationOnPurchaseBlock.NameOfElectronicPlatform,
			PurchaseLink:          dto.Dto.TabsBlock.CommonLink,
		}

		if len(dto.Dto.CustomerRequirementsBlock) > 0 {
			var participationSecurityAmount string
			if dto.Dto.CustomerRequirementsBlock[0].EnsuringPurchase.OfferGrnt {
				participationSecurityAmount = fmt.Sprintf("%d", dto.Dto.CustomerRequirementsBlock[0].EnsuringPurchase.AmountEnforcement)
			}
			purchase.ParticipationSecurityAmount = participationSecurityAmount

			if dto.Dto.CustomerRequirementsBlock[0].EnsuringPerformanceContract.OfferGrnt {
				purchase.ExecutionSecurityAmount = fmt.Sprintf("%d", dto.Dto.CustomerRequirementsBlock[0].EnsuringPerformanceContract.ContractGrntShare)
			}
		}

		lotCh <- purchase
	}
}

func upsertLot(lotCh <-chan *provider.Purchase, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	const maxUpsertLen = 3000

	defer wg.Done()

	var lot *provider.Purchase

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	m := manager.InitManager()
	defer m.Close()

	lots := make([]*provider.Purchase, 0, maxUpsertLen)

	for {
		select {
		case <-ticker.C:
			if len(lots) > 0 {
				m.UpsertPurchase(lots)

				lots = lots[:0]
			}
		case lot = <-lotCh:
			if len(lots) >= maxUpsertLen {
				m.UpsertPurchase(lots)

				lots = lots[:0]
			}

			lots = append(lots, lot)
		case <-doneCh:
			if len(lots) != 0 {
				m.UpsertPurchase(lots)
			}

			return
		}
	}
}
