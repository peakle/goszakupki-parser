package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/peakle/goszakupki-parser/internal/manager"
	"github.com/peakle/goszakupki-parser/internal/provider"
	"github.com/urfave/cli"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"
)

var m *manager.SQLManager

// StartServer - start api server
func StartServer(_ *cli.Context) {
	m = manager.InitManager()
	defer m.Close()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		path := strings.ToLower(string(ctx.Path()))

		if strings.HasPrefix(path, "/get/purchase") && string(ctx.Request.Header.Method()) == fasthttp.MethodGet {
			handlePurchase(ctx)
		} else if strings.HasPrefix(path, "/debug/pprof") {
			pprofhandler.PprofHandler(ctx)
		} else {
			ctx.SetConnectionClose()
		}
	}

	server := fasthttp.Server{
		Handler:              requestHandler,
		IdleTimeout:          30 * time.Second,
		TCPKeepalivePeriod:   provider.DefaultTimeout,
		TCPKeepalive:         true,
		MaxKeepaliveDuration: 30 * time.Second,
		ReadTimeout:          provider.DefaultTimeout,
		WriteTimeout:         provider.DefaultTimeout,
	}
	log.Fatal(server.ListenAndServe(":80"))
}

func handlePurchase(ctx *fasthttp.RequestCtx) {
	var err error
	var entryDto provider.EntryDto
	var result []provider.Purchase
	var ans []byte

	ctx.Response.Header.Set("Content-Type", "application/json")

	err = json.Unmarshal(ctx.Request.Body(), &entryDto)
	if err != nil {
		fmt.Fprint(ctx, "{\"error\": \""+err.Error()+"\"}")
		return
	}

	result, err = m.GetLots(entryDto)
	if err != nil {
		fmt.Fprint(ctx, "{\"error\": \""+err.Error()+"\"}")

		return
	}

	ans, err = json.Marshal(result)
	if err != nil {
		fmt.Fprint(ctx, "{\"error\": \""+err.Error()+"\"}")

		return
	}

	fmt.Fprint(ctx, string(ans))
}
