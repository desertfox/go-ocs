package main

import (
	"os"

	ocs "github.com/desertfox/ocs/pkg"
)

func main() {

	if len(os.Args) > 1 {
		FlagSet.Parse(os.Args[2:])
	}

	ocs := ocs.Ocs{
		Server: CLIServer,
		Token:  CLIToken,
		Config: ocs.GetOCSConfig(),
	}

	ocs.DoCommand(CLICommand)
}
