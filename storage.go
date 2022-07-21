package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

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
