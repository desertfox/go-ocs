package main

import (
	"fmt"
	"os/exec"
)

func login(c *config) {
	var (
		output []byte
		err    error
		h      host = c.getSelectedHost()
	)

	fmt.Println(yellow.Render(fmt.Sprintf("Logging into Server: %v", h.Server)))

	switch h.Version {
	case "4.X":
		output, err = exec.Command("oc", "login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server)).Output()
	case "3.X":
		output, err = exec.Command("oc", "login", h.Server, fmt.Sprintf("--token=%v", h.Token)).Output()
	default:
		output, err = exec.Command("oc", "login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server)).Output()
	}

	if err != nil {
		fmt.Println(red.Render(err.Error()))
	}

	fmt.Println(green.Render(string(output[:])))
}
