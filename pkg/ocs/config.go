package ocs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	configFileName string = ".ocsconfig"
)

type config struct {
	Selected int `yaml:"Selected"`
	Hosts    []host
}

type host struct {
	Server  string    `yaml:"Server"`
	Token   string    `yaml:"Token"`
	Created time.Time `yaml:"Created"`
}

func GetConfig() *config {
	c := EmptyConfig()

	file, err := ioutil.ReadFile(buildConfigFilePath())
	if err != nil {
		WriteConfig(c)
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return c
}

func WriteConfig(c *config) {
	data, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println("unable to marshal config data:" + err.Error())
		return
	}

	err = ioutil.WriteFile(buildConfigFilePath(), data, 0644)
	if err != nil {
		fmt.Println("unable to write data to config file" + err.Error())
		return
	}
}

func NewHost(server, token string) host {
	return host{
		Server:  server,
		Token:   token,
		Created: time.Now(),
	}
}

func buildConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("unable to find homedir" + err.Error())
	}

	return filepath.Join(home, configFileName)
}
func EmptyConfig() *config {
	return &config{
		Selected: 0,
		Hosts:    []host{},
	}
}

func (c *config) Add(h host) host {
	if c.exists(h.Server) {
		c.Update(h)

		return c.getSelectedHost()
	}

	c.Hosts = append(c.Hosts, h)
	c.setSelected(len(c.Hosts) - 1)

	printNewHost(c, h)

	return c.getSelectedHost()
}

func (c *config) Del(index int) {
	c.Hosts = append(c.Hosts[:index], c.Hosts[index+1:]...)
	c.setSelected(0)
}

func (c *config) Update(h host) {
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

func (c *config) Prune() {
	var checkTime = time.Now().Add(-1 * 24 * time.Hour)

	for i, v := range c.Hosts {
		if v.Created.After(checkTime) {
			printPuneHost(c.Hosts[i])
			c.Del(i)
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

func (c *config) CycleSelected() host {
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
