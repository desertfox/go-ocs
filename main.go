package main

import (
	_ "embed"
	"os"

	ocs "github.com/desertfox/ocs/pkg"
)

var (
	//go:embed README.md
	readme string
)

func main() {
	command, server, token, version := parseArgs(os.Args)

	ocs := ocs.New(ocs.NewHost(server, token, version))

	ocs.DoCommand(command)
}
