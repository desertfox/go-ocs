package ocs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	LEGACY bool = false
)

type Ocs struct {
	Host   host
	Config *config
}

func New(h host) Ocs {
	return Ocs{
		Host: h,
	}
}

func (o *Ocs) BuildConfig() {
	o.Config = getConfig()
}

func (o Ocs) DoCommand(CLICommand string) {
	switch CLICommand {

	case "login":
		o.add()
	case "swap":
		o.swap()
	case "list":

	case "clear":
		o.clear()
	case "cycle":
		o.cycle()
	case "del":
		o.del()
	case "prune":
		o.prune()
	default:
		o.cycle()
	}

	o.Config.writeConfig()

	o.list()
}

func (o Ocs) execLogin() {
	fmt.Println(yellow.Render(fmt.Sprintf("Logging into Server: %v", o.Host.Server)))

	if o.Host.Version == "" {
		panic("Host does not have version specified, clear your config and re-login to servers.")
	}

	var (
		output []byte
		err    error
	)

	switch o.Host.Version {
	case "4.X":
		output, err = exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Host.Token), fmt.Sprintf("--server=%v", o.Host.Server)).Output()
	case "3.X":
		output, err = exec.Command("oc", "login", o.Host.Server, fmt.Sprintf("--token=%v", o.Host.Token)).Output()
	}

	if err != nil {
		fmt.Println(red.Render(err.Error()))
	} else {
		fmt.Println(green.Render(string(output[:])))
	}

}

func (o *Ocs) add() {
	o.Host = o.Config.addHost(o.Host)

	o.execLogin()
}

func (o *Ocs) swap() {
	i, _ := strconv.Atoi(os.Args[2])

	o.Host = o.Config.swapHost(i)

	o.execLogin()
}

func (o *Ocs) del() {
	i, _ := strconv.Atoi(os.Args[2])

	o.Config.delHost(i)
}

func (o Ocs) clear() {
	o.Config.clearHost()
}

func (o *Ocs) cycle() {
	o.Host = o.Config.cycleHost()

	o.execLogin()
}

func (o Ocs) list() {
	selectedColor := strconv.Itoa(25 + o.Config.Selected*20)
	selectedString := generateStyleForeground(selectedColor).Render(fmt.Sprintf("Selected: %v ", o.Config.Selected))
	colorBar := generateStyleBackground(selectedColor).Render(strings.Repeat(" ", 10))

	fmt.Println(selectedString + colorBar + "\n")

	var colorIndex = 25
	for i, v := range o.Config.Hosts {
		hostString := generateStyleForeground(strconv.Itoa(colorIndex)).Render(fmt.Sprintf("Index:%v, Created:%v, Server:%v", i, v.Created.Format(time.RFC1123), v.Server))
		fmt.Println(hostString)
		colorIndex += 20
	}
}

func (o *Ocs) SetUpdateCheck() {
	o.Config.UpdateCheck = time.Now()
}

func (o *Ocs) prune() {
	o.Config.prune()
}
