package parser

import (
	"fmt"
	"os"
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
	var doneChan = make(chan struct{})
	defer func() {
		doneChan <- struct{}{}
		close(proxyChan)
	}()

	go proxy.LoadProxy(proxyChan, doneChan)

	loatChan := make(chan provider.Loat, 1000)
	upsertWg := &sync.WaitGroup{}

	upsertWg.Add(1)
	go upsertLoat(loatChan, upsertWg)

	//TODO implement logic

	upsertWg.Wait()
	fmt.Println("End time: ", time.Now().Format("2006-01-02 15:04"))

	return nil
}

func upsertLoat(loatCh <-chan provider.Loat, wg *sync.WaitGroup) {
	const maxUpsertLen = 10000
	defer wg.Done()

	var loat provider.Loat

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	m := manager.InitCustomManager(
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)
	defer m.Close()

	loats := make([]provider.Loat, 0, 10000)

	for {
		select {
		case <-ticker.C:
			if len(loats) > 0 {
				m.UpsertLoats(loats)

				loats = loats[:0]
			}
		default:
			if len(loats) >= maxUpsertLen {
				m.UpsertLoats(loats)

				loats = loats[:0]
			}

			loats = append(loats, loat)
		}
	}
}
