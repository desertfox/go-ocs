package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func login(c *config) error {
	var (
		h    host = c.getSelectedHost()
		args []string
		out  bytes.Buffer
	)

	fmt.Println(yellow.Render(fmt.Sprintf("Logging into Server: %v\n", h.Server)))

	switch h.Version {
	case "4.X":
		args = []string{"login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server)}
	case "3.X":
		args = []string{"login", h.Server, fmt.Sprintf("--token=%v", h.Token)}
	default:
		args = []string{"login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server)}
	}

	cmd := exec.Command("oc", args...)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		var (
			es string
			ee *exec.ExitError
		)

		if errors.As(err, &ee) {
			es = fmt.Sprintf("exit code error:%d", ee.ExitCode())
		} else {
			es = err.Error()
		}

		fmt.Println(red.Render(es))

		return errors.New(es)
	}

	fmt.Println(green.Render(out.String()))

	return nil
}
