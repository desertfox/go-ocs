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
		if err := login(c); err != nil {
			c.del(c.Selected)
		}
	case "swap":
		c.swapSelected(getOption(args))
		if err := login(c); err != nil {
			c.del(c.Selected)
		}
	case "list":
	case "clear":
		writeConfig(emptyConfig())
	case "cycle":
		c.cycleSelected()
		if err := login(c); err != nil {
			c.del(c.Selected)
		}
	case "del":
		c.del(getOption(args))
	case "prune":
		c.prune()
		if err := login(c); err != nil {
			c.del(c.Selected)
		}
	default:
		c.cycleSelected()
		if err := login(c); err != nil {
			c.del(c.Selected)
		}
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
