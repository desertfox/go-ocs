package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	configFileName string = ".ocsconfig"
)

func buildConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("unable to find homedir" + err.Error())
	}

	return filepath.Join(home, configFileName)
}

type config struct {
	Selected int `yaml:"Selected"`
	Hosts    []host
}

func emptyConfig() *config {
	return &config{
		Selected: 0,
		Hosts:    []host{},
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
	c.setSelected(0)
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

func (c *config) setSelected(i int) {
	c.Selected = i
}

func (c *config) getSelectedHost() host {
	return c.Hosts[c.Selected]
}

func (c *config) swapSelected(index int) host {
	if index > len(c.Hosts)-1 {
		fmt.Printf("Swap %v greater than %v of config values", index, len(c.Hosts))
	} else {
		c.setSelected(index)
	}

	return c.getSelectedHost()
}

func (c *config) cycleSelected() host {
	if len(c.Hosts) <= 1 {
		fmt.Printf("%v Host configured, no-op.\n", len(c.Hosts))
		os.Exit(0)
	}

	if c.Selected+1 > len(c.Hosts)-1 {
		c.Selected = 0
	} else {
		c.Selected++
	}

	return c.getSelectedHost()
}
