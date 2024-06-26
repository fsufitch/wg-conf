package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	generateCommand := InitializeCommand{}
	clientAddCommand := ClientAddCommand{}
	printIniCommand := PrintIniCommand{}

	var app = &cli.App{
		Name:        "wg-conf",
		Description: "An opinionated tool for managing Wireguard configuration files",
		Authors:     []*cli.Author{{Name: "Filip Sufitchi", Email: "fsufitchi@gmail.com"}},
		Suggest:     true,
		Commands: []*cli.Command{
			generateCommand.Command(),
			clientAddCommand.Command(),
			printIniCommand.Command(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(fmt.Sprintf("unexpected error returned from app: %v", err))
	}
}
