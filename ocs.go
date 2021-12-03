package main

import (
	"os"

	ocs "github.com/desertfox/ocs/pkg"
	"github.com/desertfox/ocs/pkg/config"
)

func main() {

	if len(os.Args) > 1 {
		FlagSet.Parse(os.Args[2:])
	}

	host := config.Host{
		Server: CLIServer,
		Token:  CLIToken,
	}

	config := config.GetOCSConfig()

	ocs := ocs.Ocs{
		Host:   host,
		Config: config,
	}

	ocs.DoCommand(CLICommand)
}
