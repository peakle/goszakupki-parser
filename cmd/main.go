package main

import (
	"fmt"
	"os"

	"github.com/peakle/goszakupki-parser/pkg/parser"
	"github.com/peakle/goszakupki-parser/pkg/server"
	"github.com/urfave/cli"
)

var (
	// Version - app release
	Version = "0"
	// CommitID - release's commid id
	CommitID = "0"
	commands = []cli.Command{
		{
			Name:        "parse-lot-44",
			ShortName:   "44-fz",
			Description: "parse lots for 44-fz",
			Action:      parser.ProcessLot44,
			Category:    "parser",
		},
		{
			Name:        "lots-server",
			ShortName:   "server",
			Description: "give lots data",
			Action:      server.StartServer,
			Category:    "server",
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "lot"
	app.Commands = commands
	app.Version = fmt.Sprintf("%s - %s", Version, CommitID)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
}
