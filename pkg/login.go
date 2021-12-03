package ocs

import (
	"fmt"
	"os/exec"
)

func (o Ocs) add() {
	o.execLogin()

	o.Config.add(o.Server, o.Token)
}

func (o Ocs) execLogin() {
	fmt.Printf("Logging into Server: %v\n", o.Server)

	exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Token), fmt.Sprintf("--server=%v", o.Server))
}
