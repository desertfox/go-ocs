package main

import (
	_ "embed"
	"os"
)

var (
	//go:embed README.md
	readme string
)

func main() {
	command, server, token, version := parseArgs(os.Args)

	ocs := Ocs{
		host:   newHost(server, token, version),
		config: getConfig(),
	}

	ocs.DoCommand(command)
}
