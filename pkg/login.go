package ocs

import (
	"fmt"
	"os/exec"
)

func (o Ocs) add() {
	o.execLogin()

	o.Config.addHost(o.Host)
}

func (o Ocs) execLogin() {
	fmt.Printf("Logging into Server: %v\n", o.Host.Server)

	exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Host.Token), fmt.Sprintf("--server=%v", o.Host.Server))
}
