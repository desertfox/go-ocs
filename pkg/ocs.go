package ocs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

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
	fmt.Printf("Selected: %v\n", o.Config.Selected)

	for i, v := range o.Config.Hosts {
		fmt.Printf("%v:%#v\n", i, v)
	}
}

func (o Ocs) add() {
	o.execLogin()

	o.Config.AddHost(o.Host)
}

func (o Ocs) execLogin() {
	fmt.Printf("Logging into Server: %v\n", o.Host.Server)

	output, err := exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Host.Token), fmt.Sprintf("--server=%v", o.Host.Server)).Output()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(output[:]))
	}

}
func (o Ocs) clear() {
	o.Config.Clear()
}
func (o *Ocs) swap() {

	id := os.Args[2]

	i, err := strconv.Atoi(id)
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
