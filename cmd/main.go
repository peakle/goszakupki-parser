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
			Name:        "parse-loat-44",
			ShortName:   "44-fz",
			Description: "parse loats for 44-fz",
			Action:      parser.ProcessLoat44,
			Category:    "parser",
		},
		{
			Name:        "loats-server",
			ShortName:   "server",
			Description: "give loats data",
			Action:      server.StartServer,
			Category:    "server",
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "loats"
	app.Commands = commands
	app.Version = fmt.Sprintf("%s - %s", Version, CommitID)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
}
