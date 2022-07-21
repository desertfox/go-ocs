package main

import (
	"fmt"
	"strconv"
)

func Do(CLICommand string, h host, args []string) {
	var c *config = getConfig()

	switch CLICommand {
	case "login":
		c.add(h)
		c.login()
	case "swap":
		c.swapSelected(getOption(args))
		c.login()
	case "list":
	case "clear":
		writeConfig(emptyConfig())
	case "cycle":
		c.cycleSelected()
		c.login()
	case "del":
		c.del(getOption(args))
		c.login()
	case "prune":
		c.prune()
		c.login()
	default:
		c.cycleSelected()
		c.login()
	}

	writeConfig(c)

	printStatus(c)
}

func getOption(args []string) int {
	if len(args) <= 2 {
		fmt.Print("Missing Argument")
	}

	i, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Print(err)
	}
	return i
}
