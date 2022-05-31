package main

import (
	"os"
)

func main() {
	command, server, token, version := parseArgs(os.Args)

	Do(command, newHost(server, token, version), os.Args)
}
