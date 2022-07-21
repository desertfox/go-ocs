package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type host struct {
	Server  string    `yaml:"Server"`
	Token   string    `yaml:"Token"`
	Version string    `yaml:"Version"`
	Created time.Time `yaml:"Created"`
}

func newHost(server, token, version string) host {
	return host{
		Server:  server,
		Token:   token,
		Version: version,
		Created: time.Now(),
	}
}

func (c *config) add(h host) host {
	if c.exists(h.Server) {
		c.update(h)
		return c.getSelectedHost()
	}

	c.Hosts = append(c.Hosts, h)
	c.setSelected(len(c.Hosts) - 1)

	printNewHost(c, h)

	return c.getSelectedHost()
}

func (c *config) del(index int) {
	c.Hosts = append(c.Hosts[:index], c.Hosts[index+1:]...)
}

func (c *config) update(h host) {
	for i, host := range c.Hosts {
		if host.Server == h.Server {
			c.Hosts[i] = h
			c.setSelected(i)
			printUpdateHost(h)
			break
		}
	}
}

func (c *config) exists(server string) bool {
	for _, host := range c.Hosts {
		if host.Server == server {
			printServerExists(server)
			return true
		}
	}
	return false
}

func (c *config) prune() {
	var checkTime = time.Now().Add(-1 * 24 * time.Hour)

	for i, v := range c.Hosts {
		if v.Created.After(checkTime) {
			printPuneHost(c.Hosts[i])
			c.del(i)
			if c.Selected == i {
				c.setSelected(0)
			}
			continue
		}
	}
}

func (c *config) login() {
	var (
		h    host = c.getSelectedHost()
		args []string
		ee   *exec.ExitError
		out  bytes.Buffer
	)

	fmt.Println(yellow.Render(fmt.Sprintf("Logging into Server: %v", h.Server)))

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
		if errors.As(err, &ee) {
			fmt.Println(red.Render(fmt.Sprintf("exit code error:%d", ee.ExitCode())))
		} else {
			fmt.Println(red.Render(err.Error()))
		}
		return
	}

	fmt.Println(green.Render(out.String()))
}
