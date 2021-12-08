package ocs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

var configFile string = ".ocsconfig"

type config struct {
	Selected    int       `yaml:"Selected"`
	UpdateCheck time.Time `yaml:UpdateCheck"`
	Hosts       []Host
}

type Host struct {
	Server  string    `yaml:"Server"`
	Token   string    `yaml:"Token"`
	Created time.Time `yaml:"Created"`
}

func (c *config) SetSelected(i int) {
	c.Selected = i
}

func GetConfig() *config {
	c := &config{
		Selected:    0,
		Hosts:       []Host{},
		UpdateCheck: time.Now(),
	}

	file, err := ioutil.ReadFile(c.getConfigFilePath())
	if err != nil {
		c.writeConfig()
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return c
}

func (c *config) writeConfig() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	err = ioutil.WriteFile(c.getConfigFilePath(), data, 0644)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}
}

func (c *config) addHost(h Host) Host {
	if c.serverExists(h.Server) {
		c.updateHost(h)
		return c.GetSelectedHost()
	}

	c.Hosts = append(c.Hosts, h)
	c.SetSelected(len(c.Hosts) - 1)

	selectedColor := lipgloss.Color(strconv.Itoa(25 + c.Selected*20))
	addHostString := style.PaddingLeft(2).Foreground(selectedColor).Render(fmt.Sprintf("AddHost: %v\n", h.Server))
	fmt.Println(addHostString)

	return c.GetSelectedHost()
}

func (c *config) swapHost(index int) Host {
	if index > len(c.Hosts)-1 {
		fmt.Printf("Swap %v greater than %v of config values", index, len(c.Hosts)-1)
	} else {
		c.SetSelected(index)
	}

	return c.GetSelectedHost()
}

func (c *config) delHost(index int) {
	c.Hosts = append(c.Hosts[:index], c.Hosts[index+1:]...)
}

func (c *config) cycleHost() Host {
	if len(c.Hosts) <= 1 {
		fmt.Printf("%v Host configured, no-op.\n", len(c.Hosts))
	}

	if c.Selected+1 > len(c.Hosts)-1 {
		c.Selected = 0
	} else {
		c.Selected++
	}

	return c.GetSelectedHost()
}

func (c config) serverExists(server string) bool {
	for _, host := range c.Hosts {
		if host.Server == server {
			serverExistsString := style.PaddingLeft(2).Foreground(lipgloss.Color("11")).Render(fmt.Sprintf("serverExists: %v", server))
			fmt.Println(serverExistsString)
			return true
		}
	}
	return false
}

func (c *config) updateHost(h Host) {
	for i, host := range c.Hosts {
		if host.Server == h.Server {
			c.Hosts[i] = h
			c.SetSelected(i)

			updateHostString := style.PaddingLeft(2).Foreground(lipgloss.Color("10")).Render(fmt.Sprintf("updateHost: %v\n", h.Server))
			fmt.Println(updateHostString)

			break
		}
	}
}

func (c config) getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return filepath.Join(home, configFile)
}

func (c config) GetSelectedHost() Host {
	return c.Hosts[c.Selected]
}

func (c *config) clearHost() {
	c.Hosts = []Host{}
	c.Selected = 0
}
