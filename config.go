package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
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

func getConfig() *config {
	c := emptyConfig()

	file, err := ioutil.ReadFile(buildConfigFilePath())
	if err != nil {
		writeConfig(c)
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return c
}

func writeConfig(c *config) {
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
