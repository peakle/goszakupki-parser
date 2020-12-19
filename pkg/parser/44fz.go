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

	lotChan := make(chan provider.Lot, 1000)
	upsertWg := &sync.WaitGroup{}

	upsertWg.Add(1)
	go upsertLot(lotChan, doneChan, upsertWg)

	concurCh := make(chan struct{}, 10) // increase for more parallelism
	lot44Logic(lotChan, <-proxyChan, concurCh)

	upsertWg.Wait()
	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
}

func lot44Logic(lotCh chan<- provider.Lot, proxy string, concurCh chan struct{}) {
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

	err = client.DoTimeout(req, resp, provider.DefaultTimeout)
	if err != nil {
		log.Println("on lot44Logic: on DoTimeout: " + err.Error())
		return
	}

	var dto *provider.Dto44fz
	err = json.Unmarshal(resp.Body(), dto)
	if dto == nil && err != nil {
		log.Println("on lot44Logic: on unmarshal: " + err.Error())
		return
	}

}

func upsertLot(lotCh <-chan provider.Lot, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	const maxUpsertLen = 10000

	defer wg.Done()

	var lot provider.Lot

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	m := manager.InitManager()
	defer m.Close()

	lots := make([]provider.Lot, 0, 10000)

	for {
		select {
		case <-ticker.C:
			if len(lots) > 0 {
				m.UpsertLots(lots)

				lots = lots[:0]
			}
		case lot = <-lotCh:
			if len(lots) >= maxUpsertLen {
				m.UpsertLots(lots)

				lots = lots[:0]
			}

			lots = append(lots, lot)
		case <-doneCh:
			if len(lots) != 0 {
				m.UpsertLots(lots)
			}

			return
		}
	}
}
