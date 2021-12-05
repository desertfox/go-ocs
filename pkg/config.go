package ocs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ocsconfigFile string = ".ocsconfig"

type Ocsconfig struct {
	Selected int `yaml:"Selected"`
	Hosts    []Host
}

type Host struct {
	Server string `yaml:"Server"`
	Token  string `yaml:"Token"`
}

func GetOCSConfig() *Ocsconfig {
	oc := &Ocsconfig{
		Selected: 0,
		Hosts:    []Host{},
	}

	file, err := ioutil.ReadFile(oc.getConfigFilePath())
	if err != nil {
		oc.WriteConfig()
	}

	err = yaml.Unmarshal(file, &oc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return oc
}

func (oc *Ocsconfig) WriteConfig() {
	data, err := yaml.Marshal(&oc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	err = ioutil.WriteFile(oc.getConfigFilePath(), data, 0644)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}
}

func (oc *Ocsconfig) AddHost(h Host) {
	if oc.serverExists(h.Server) {
		oc.updateHost(h)

		fmt.Printf("AddHost: server exists, updating server %v with new token\n", h.Server)

		return
	}

	oc.Hosts = append(oc.Hosts, h)

	fmt.Printf("AddHost: %v\n", h.Server)

}

func (oc Ocsconfig) serverExists(server string) bool {
	for _, host := range oc.Hosts {
		if host.Server == server {
			return true
		}
	}
	return false
}

func (oc *Ocsconfig) SetSelected(i int) {
	oc.Selected = i
}

func (oc *Ocsconfig) updateHost(h Host) {
	for i, host := range oc.Hosts {
		if host.Server == h.Server {
			oc.Hosts[i] = h
			break
		}
	}
}

func (oc Ocsconfig) getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return filepath.Join(home, ocsconfigFile)
}

func (oc Ocsconfig) GetSelectedHost() Host {
	return oc.Hosts[oc.Selected]
}

func (oc *Ocsconfig) Clear() {
	oc.Hosts = []Host{}
	oc.Selected = 0
}
