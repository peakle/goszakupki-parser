package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/peakle/goszakupki-parser/pkg/manager"
	"github.com/peakle/goszakupki-parser/pkg/provider"
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

		if strings.HasPrefix(path, "/get/loat") {
			handle(ctx)
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

func handle(ctx *fasthttp.RequestCtx) {
	var err error
	var entryDto provider.EntryDto
	var result provider.Result

	ctx.Response.Header.Set("Content-Type", "application/json")

	err = json.Unmarshal(ctx.Request.Body(), &entryDto)
	if err != nil {
		fmt.Fprint(ctx, "{\"error\": \""+err.Error()+"\"}")
		return
	}

	result, err = m.GetLoats(entryDto)
	if err != nil {
		fmt.Fprint(ctx, "{\"error\": \""+err.Error()+"\"}")

		return
	}

	ans := bytes.NewBuffer()
	ans, err = json.Marshal(result)
	if err != nil {
		fmt.Fprintf(ctx, "{\"error\": \""+err.Error()+"\"}")

		return
	}

	fmt.Fprint(ctx, ans.String())
}
