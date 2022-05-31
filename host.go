package main

import (
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
			continue
		}
	}
}
