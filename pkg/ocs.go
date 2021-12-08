package ocs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().PaddingLeft(2)

type Ocs struct {
	Host   Host
	Config *config
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
	default:
		o.cycle()
	}

	o.Config.writeConfig()

	o.list()
}

func (o Ocs) execLogin() {
	loginString := style.Foreground(lipgloss.Color("11")).Render(fmt.Sprintf("Logging into Server: %v", o.Host.Server))
	fmt.Println(loginString)

	output, err := exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Host.Token), fmt.Sprintf("--server=%v", o.Host.Server)).Output()

	if err != nil {
		fmt.Println(style.Foreground(lipgloss.Color("9")).Render(err.Error()))
	} else {
		fmt.Println(style.Foreground(lipgloss.Color("10")).Render(string(output[:])))
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
	selectedString := style.Foreground(lipgloss.Color(selectedColor)).Render(fmt.Sprintf("Selected: %v ", o.Config.Selected))
	colorBar := style.Copy().Padding(0, 0, 0, 0).Background(lipgloss.Color(selectedColor)).Render(strings.Repeat(" ", 10))

	fmt.Println(selectedString + colorBar + "\n")

	var colorIndex = 25
	for i, v := range o.Config.Hosts {
		hostString := style.Padding(0, 2, 0, 2).Foreground(lipgloss.Color(strconv.Itoa(colorIndex))).Render(fmt.Sprintf("Index:%v, Created:%v, Server:%v", i, v.Created.Format(time.RFC1123), v.Server))
		fmt.Println(hostString)
		colorIndex += 20
	}
}

func (o *Ocs) SetUpdateCheck() {
	o.Config.UpdateCheck = time.Now()
}
