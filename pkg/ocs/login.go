package ocs

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func Login(c *config) error {
	var (
		h   host = c.getSelectedHost()
		out bytes.Buffer
	)

	printColor("yellow", fmt.Sprintf("Logging into Server: %v\n", h.Server))

	cmd := exec.Command("oc", "login", fmt.Sprintf("--token=%v", h.Token), fmt.Sprintf("--server=%v", h.Server))
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

		printColor("red", es)

		return errors.New(es)
	}

	printColor("green", out.String())

	return nil
}
