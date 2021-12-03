package ocs

import (
	"fmt"
	"os/exec"
)

func (o Ocs) add() {

	o.execLogin()

	o.config.add(o.Server, o.Token)
}

func (o Ocs) execLogin() {
	fmt.Printf("Logging into Server: %v", o.Server)
	exec.Command("oc", "login", fmt.Sprintf("--token=%v", o.Token), fmt.Sprintf("--server=%v", o.Server))
}
