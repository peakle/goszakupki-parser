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

// ProcessLoat44 collect data about new loats for 44-FZ
func ProcessLoat44(_ *cli.Context) error {
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

	loatChan := make(chan provider.Lot, 1000)
	upsertWg := &sync.WaitGroup{}

	upsertWg.Add(1)
	go upsertLoat(loatChan, doneChan, upsertWg)

	//TODO implement logic

	upsertWg.Wait()
	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
}

func upsertLoat(loatCh <-chan provider.Lot, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	const maxUpsertLen = 10000

	defer wg.Done()

	var loat provider.Lot

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	m := manager.InitManager()
	defer m.Close()

	loats := make([]provider.Lot, 0, 10000)

	for {
		select {
		case <-ticker.C:
			if len(loats) > 0 {
				m.UpsertLoats(loats)

				loats = loats[:0]
			}
		case loat = <-loatCh:
			if len(loats) >= maxUpsertLen {
				m.UpsertLoats(loats)

				loats = loats[:0]
			}

			loats = append(loats, loat)
		case <-doneCh:
			return
		}
	}
}
