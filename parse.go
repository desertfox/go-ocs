package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

var (
	//go:embed README.md
	readme string
)

func parseArgs(args []string) (string, string, string, string) {
	var (
		CLICommand, CLIToken, CLIServer string
		CLIVersion                      string = "4.X"
		FlagSet                         flag.FlagSet
	)

	FlagSet = flag.FlagSet{}
	FlagSet.StringVar(&CLIServer, "server", "", "Server")
	FlagSet.StringVar(&CLIToken, "token", "", "Token")

	if len(args) <= 1 {
		return "", "", "", ""
	}

	args = args[1:]

	CLICommand = args[0]
	if CLICommand == "help" || CLICommand == "h" {
		fmt.Println(readme)
		os.Exit(0)
	}

	if len(args) == 1 {
		return CLICommand, "", "", ""
	}

	if CLICommand == "oc" {
		args = args[1:]
		CLICommand = args[0]

		FlagSet.Parse(args[1:])

		if CLIServer == "" && CLIToken == "" {
			FlagSet.Parse(args[2:])
			CLIServer = args[1]
			CLIVersion = "3.X"
		}

		return CLICommand, CLIServer, CLIToken, CLIVersion
	}

	return CLICommand, "", "", ""
}
