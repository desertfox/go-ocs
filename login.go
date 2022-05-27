package main

import (
	"fmt"
	"os/exec"
)

func login(h host) {
	fmt.Println(yellow.Render(fmt.Sprintf("Logging into Server: %v", h.Server)))

	if h.Version == "" {
		panic("Host does not have version specified, clear your config and re-login to servers.")
	}

	var (
		output []byte
		err    error
	)

	switch h.Version {
	case "4.X":
		output, err = exec.Command("oc", "login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server)).Output()
	case "3.X":
		output, err = exec.Command("oc", "login", h.Server, fmt.Sprintf("--token=%v", h.Token)).Output()
	}

	if err != nil {
		fmt.Println(red.Render(err.Error()))
	} else {
		fmt.Println(green.Render(string(output[:])))
	}

}
