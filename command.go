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
	case "swap":
		c.swapSelected(getOption(args))
	case "list":
	case "clear":
		writeConfig(emptyConfig())
		return
	case "cycle":
		c.CycleSelected()
	case "del":
		c.del(getOption(args))
	case "prune":
		c.prune()
		return
	default:
		c.CycleSelected()
	}

	login(c)

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
