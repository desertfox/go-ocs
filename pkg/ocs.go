package ocs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().PaddingLeft(2)

type Ocs struct {
	Host   Host
	Config *Ocsconfig
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

	o.Config.WriteConfig()

	o.list()
}

func (o *Ocs) cycle() {
	if len(o.Config.Hosts) <= 1 {
		fmt.Printf("%v Host configured, no-op.\n", len(o.Config.Hosts))
		return
	}

	selected := o.Config.Selected

	if selected+1 > len(o.Config.Hosts)-1 {
		o.Config.Selected = 0
	} else {
		o.Config.Selected++
	}

	o.Host = o.Config.GetSelectedHost()

	o.execLogin()
}

func (o Ocs) list() {
	selectedColor := strconv.Itoa(25 + o.Config.Selected*20)
	selectedString := style.Foreground(lipgloss.Color(selectedColor)).Render(fmt.Sprintf("Selected: %v ", o.Config.Selected))
	colorBar := style.Copy().Padding(0, 0, 0, 0).Background(lipgloss.Color(selectedColor)).Render(strings.Repeat(" ", 10))
	fmt.Println(selectedString + colorBar + "\n")

	var colorIndex = 25
	for i, v := range o.Config.Hosts {
		hostString := style.Padding(0, 2, 0, 2).Foreground(lipgloss.Color(strconv.Itoa(colorIndex))).Render(fmt.Sprintf("Index:%v, Server:%v Created:%v", i, v.Server, v.Created.String()))
		fmt.Println(hostString)
		colorIndex += 20
	}
}

func (o Ocs) add() {
	o.execLogin()

	o.Config.AddHost(o.Host)
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
func (o Ocs) clear() {
	o.Config.Clear()
}
func (o *Ocs) swap() {
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	if i > len(o.Config.Hosts)-1 {
		fmt.Printf("Swap %v greater than %v of config values", i, len(o.Config.Hosts)-1)
		return
	} else {
		o.Config.SetSelected(i)
	}

	o.Host = o.Config.GetSelectedHost()

	o.execLogin()
}

func (o *Ocs) del() {
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	o.Config.Hosts = append(o.Config.Hosts[:i], o.Config.Hosts[i+1:]...)
}
