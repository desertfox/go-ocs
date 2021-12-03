package main

import (
	"os"

	ocs "github.com/desertfox/ocs/pkg"
)

func main() {

	if len(os.Args) > 1 {
		FlagSet.Parse(os.Args[2:])
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
