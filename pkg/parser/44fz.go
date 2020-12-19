package parser

import (
	"crypto/tls"
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
func ProcessLot44(c *cli.Context) error {
	fmt.Println("Start time: ", time.Now().Format("2006-01-02 15:04"))

	fromDate := c.String("from-date") // may be add iterate through dates slice
	toDate := c.String("to-date")

	if fromDate == "" {
		fromDate = time.Now().Format("02-01-2006")
	}

	if toDate == "" {
		toDate = time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	}

	var regNumber string
	var proxyChan = make(chan string, 3000)
	var doneChan = make(chan struct{}, 2)
	var lotChan = make(chan *provider.Purchase, 1000)
	var concurCh = make(chan struct{}, 1) // increase for more parallelism
	var regNumberCh = make(chan string, 10000)

	var workerWg = &sync.WaitGroup{}
	var upsertWg = &sync.WaitGroup{}

	upsertWg.Add(1)

	defer func() {
		close(doneChan)
		close(proxyChan)
	}()

	go proxy.LoadProxy(proxyChan, doneChan)
	go insertLot(lotChan, doneChan, upsertWg)
	go fz44RegNumberGenerator(fromDate, toDate, regNumberCh, proxyChan)

	for regNumber = range regNumberCh {
		workerWg.Add(1)
		go fz44LotWorker(regNumber, lotChan, <-proxyChan, concurCh, workerWg)
	}

	doneChan <- struct{}{} // for upserts
	doneChan <- struct{}{} // for proxy
	upsertWg.Wait()

	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
}

func fz44RegNumberGenerator(fromDate, toDate string, regNumberCh chan<- string, proxyCh <-chan string) {
	//pageNumber, from, to, price (from)
	const SearchURIPattern = "https://zakupki.gov.ru/epz/order/extendedsearch/results.html?morphology=on&sortDirection=true&recordsPerPage=_50&showLotsInfoHidden=false&sortBy=PRICE&fz44=on&af=on&ca=on&currencyIdGeneral=-1&formatInJson=true&pageNumber=%d&publishDateFrom=%s&publishDateTo=%s&priceFromGeneral=%d"
	const maxRecordsPerPage = 50
	const maxPageCount = 20

	var price int
	var err error
	var uri, proxy string
	var searchDto provider.ExtentendedSearch
	var l provider.List

	var req *fasthttp.Request
	var resp *fasthttp.Response

	var pageNumber = 1
	var client = fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	for {
		req = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()

		if proxy != "" {
			client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(proxy, 30*time.Second)
		} else {
			client.Dial = nil
		}

		uri = fmt.Sprintf(SearchURIPattern, pageNumber, fromDate, toDate, price)

		req.SetRequestURI(uri)
		req.Header.SetUserAgent(provider.UserAgent)
		req.SetConnectionClose()

		err = client.DoTimeout(req, resp, provider.DefaultTimeout)
		if err != nil {
			log.Println("on fz44regNumberGenerator: on DoTimeout:" + err.Error())

			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)

			continue
		}

		err = json.Unmarshal(resp.Body(), &searchDto)
		if err != nil {
			log.Printf("on fz44regNumberGenerator: on Unmarshal: %s, uri: %s \n", err.Error(), uri)

			if searchDto.List == nil {
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(resp)

				continue
			}
		}

		if searchDto.Total == 0 {
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)

			break
		}

		for _, l = range searchDto.List {
			if l.Number != "" {
				regNumberCh <- l.Number
			}
		}

		if searchDto.Total < maxRecordsPerPage {
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)

			break
		}

		if pageNumber == maxPageCount {
			price = int(l.Price)
			pageNumber = 1

			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)

			continue
		}

		pageNumber++
	}
}

func fz44LotWorker(regNumber string, lotCh chan<- *provider.Purchase, proxy string, concurCh chan struct{}, wg *sync.WaitGroup) {
	concurCh <- struct{}{}
	defer func() {
		<-concurCh
		wg.Done()
	}()

	var err error
	var dto provider.Dto44fz
	var uri string

	client := fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if proxy != "" {
		client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(proxy, provider.DefaultTimeout)
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	uri = fmt.Sprintf(provider.URIPatternFZ44Purchase, regNumber)

	req.SetRequestURI(uri)
	req.Header.SetUserAgent(provider.UserAgent)
	req.SetConnectionClose()

	err = client.DoTimeout(req, resp, provider.DefaultTimeout)
	if err != nil {
		log.Println("on lot44Logic: on DoTimeout: " + err.Error())
		return
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return
	}

	err = json.Unmarshal(resp.Body(), &dto)
	if err != nil {
		log.Println("on lot44Logic: on unmarshal: " + err.Error())

		if dto.Dto.HeaderBlock.PurchaseNumber == "" {
			return
		}
	}

	{
		purchase := &provider.Purchase{
			ID:           dto.Dto.HeaderBlock.PurchaseNumber,
			Fz:           dto.Dto.HeaderBlock.PlacingWayFZ,
			Customer:     dto.Dto.HeaderBlock.OrganizationPublishName,
			CustomerLink: dto.Dto.HeaderBlock.OrganizationPublishLink,
			// CustomerInn: dto.Dto.Inn,
			CustomerRegion: dto.Dto.OrganizationDefinitionSupplierBlock.Location,
			// BiddingRegion: ,
			// CustomerActivityField: dto.Dto.HeaderBlock.PurchaseObjectName,
			BiddingVolume: fmt.Sprintf("%.3f", dto.Dto.InitialContractPriceBlock.InitialContractPrice),
			// BiddingCount: ,
			PurchaseTarget:        dto.Dto.HeaderBlock.PurchaseObjectName,
			RegistryBiddingNumber: dto.Dto.HeaderBlock.PurchaseNumber,
			ContractPrice:         fmt.Sprintf("%.3f", dto.Dto.InitialContractPriceBlock.InitialContractPrice),
			PublishedAt:           time.Unix(dto.Dto.ProcedurePurchaseBlock.StartDateTime/1000, 0).Format("02-01-2006"),
			RequisitionDeadlineAt: time.Unix(dto.Dto.ProcedurePurchaseBlock.EndDateTime/1000, 0).Format("02-01-2006"),
			ContractStartAt:       "", //TODO
			ContractEndAt:         "", //TODO
			Playground:            dto.Dto.GeneralInformationOnPurchaseBlock.NameOfElectronicPlatform,
			PurchaseLink:          dto.Dto.TabsBlock.CommonLink,
		}

		if len(dto.Dto.CustomerRequirementsBlock) > 0 {
			var participationSecurityAmount string
			if dto.Dto.CustomerRequirementsBlock[0].EnsuringPurchase.OfferGrnt {
				participationSecurityAmount = fmt.Sprintf("%.3f", dto.Dto.CustomerRequirementsBlock[0].EnsuringPurchase.AmountEnforcement)
			}
			purchase.ParticipationSecurityAmount = participationSecurityAmount

			if dto.Dto.CustomerRequirementsBlock[0].EnsuringPerformanceContract.OfferGrnt {
				purchase.ExecutionSecurityAmount = fmt.Sprintf("%.3f", dto.Dto.CustomerRequirementsBlock[0].EnsuringPerformanceContract.ContractGrntShare)
			}
		}

		lotCh <- purchase
	}
}

func insertLot(lotCh <-chan *provider.Purchase, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	const maxUpsertLen = 3000

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
				m.InsertPurchase(lots)

				lots = lots[:0]
			}
		case lot = <-lotCh:
			if len(lots) >= maxUpsertLen {
				m.InsertPurchase(lots)

				lots = lots[:0]
			}

			lots = append(lots, lot)
		case <-doneCh:
			if len(lots) != 0 {
				m.InsertPurchase(lots)
			}

			return
		}
	}
}
