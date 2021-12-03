package ocs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ocsconfigFile string = ".ocsconfig"

type ocsconfig struct {
	Selected int `yaml:"Selected"`
	List     []Host
}

type Host struct {
	Server string `yaml:"Server"`
	Token  string `yaml:"Token"`
}

func GetOCSConfig() *ocsconfig {
	oc := &ocsconfig{
		Selected: 0,
		List:     []Host{},
	}

	file, err := ioutil.ReadFile(oc.getConfigFilePath())
	if err != nil {
		oc.writeConfig()
	}

	err = yaml.Unmarshal(file, &oc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return oc
}

func (oc *ocsconfig) writeConfig() {
	data, err := yaml.Marshal(&oc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	err = ioutil.WriteFile(oc.getConfigFilePath(), data, 0777)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}
}

func (oc *ocsconfig) addHost(h Host) {
	if oc.serverExists(h.Server) {
		oc.updateHost(h)

		fmt.Printf("add, server exists updating %v:", oc.List)

		return
	}

	oc.List = append(oc.List, h)

	fmt.Printf("add: %v", oc.List)

}

func (oc ocsconfig) serverExists(server string) bool {
	for _, host := range oc.List {
		if host.Server == server {
			return true
		}
	}
	return false
}

func (oc *ocsconfig) setSelected(i int) {
	oc.Selected = i
}

func (oc *ocsconfig) updateHost(h Host) {
	for i, host := range oc.List {
		if host.Server == h.Server {
			oc.List[i] = h
			break
		}
	}
}

func (oc ocsconfig) getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	return filepath.Join(home, ocsconfigFile)
}

func (oc ocsconfig) GetSelectedHost() Host {
	return oc.List[oc.Selected]
}
