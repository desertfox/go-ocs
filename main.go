package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	ocs "github.com/desertfox/ocs/pkg"
)

var (
	CLICommand, CLIToken, CLIServer string
	FlagSet                         flag.FlagSet
)

var (
	//go:embed README.md
	readme string
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

func main() {
	if len(os.Args) > 1 {
		FlagSet.Parse(os.Args[2:])

		if os.Args[1] == "help" || os.Args[1] == "h" {
			fmt.Println(readme)
			os.Exit(0)
		}
	}

	host := ocs.Host{
		Server: CLIServer,
		Token:  CLIToken,
	}
	config := ocs.GetOCSConfig()

	ocs := ocs.Ocs{
		Host:   host,
		Config: config,
	}

	ocs.DoCommand(CLICommand)
}
