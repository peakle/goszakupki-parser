package parser

import (
	"fmt"
	"sync"
	"time"

	"github.com/peakle/goszakupki-parser/pkg/manager"
	"github.com/peakle/goszakupki-parser/pkg/provider"
	"github.com/peakle/goszakupki-parser/pkg/proxy"
	"github.com/urfave/cli"
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

	//TODO implement logic

	upsertWg.Wait()
	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
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
			return
		}
	}
}
