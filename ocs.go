package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ocs struct {
	host   host
	config *config
}

func (o *Ocs) BuildConfig() {
	o.config = getConfig()
}

func (o Ocs) DoCommand(CLICommand string) {
	switch CLICommand {

	case "login":
		login(o.config.addHost(o.host))
	case "swap":
		i, _ := strconv.Atoi(os.Args[2])

		login(o.config.swapHost(i))
	case "list":

	case "clear":
		o.config.clearHost()
	case "cycle":
		login(o.config.cycleHost())
	case "del":
		i, _ := strconv.Atoi(os.Args[2])

		o.config.delHost(i)
	case "prune":
		o.config.prune()
	default:
		login(o.config.cycleHost())
	}

	o.config.writeConfig()

	selectedColor := strconv.Itoa(25 + o.config.Selected*20)
	selectedString := generateStyleForeground(selectedColor).Render(fmt.Sprintf("Selected: %v ", o.config.Selected))
	colorBar := generateStyleBackground(selectedColor).Render(strings.Repeat(" ", 10))

	fmt.Println(selectedString + colorBar + "\n")

	var colorIndex = 25
	for i, v := range o.config.Hosts {
		hostString := generateStyleForeground(strconv.Itoa(colorIndex)).Render(fmt.Sprintf("Index:%v, Created:%v, Server:%v", i, v.Created.Format(time.RFC1123), v.Server))
		fmt.Println(hostString)
		colorIndex += 20
	}

}
