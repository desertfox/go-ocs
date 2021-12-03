package main

import (
	"flag"
	"os"
)

var (
	CLICommand, CLIToken, CLIServer string
	FlagSet                         flag.FlagSet
)

func init() {

	if len(os.Args) > 1 {
		CLICommand = os.Args[1]
	} else {
		CLICommand = ""
	}

	FlagSet = flag.FlagSet{}

	FlagSet.StringVar(&CLIServer, "server", "", "Server")
	FlagSet.StringVar(&CLIToken, "token", "", "Token")

}
