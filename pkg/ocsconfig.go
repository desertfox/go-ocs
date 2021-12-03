package ocs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var ocsconfigFile string = ".ocsconfig"

type ocsconfig struct {
	Selected int `yaml:"Selected"`
	List     []string
}

func GetOCSConfig() *ocsconfig {
	oc := &ocsconfig{
		Selected: 0,
		List:     []string{},
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

func (oc *ocsconfig) add(server, token string) {
	if oc.serverExists(server) {
		oc.updateHost(server, token)

		fmt.Printf("add, server exists updating %v:", oc.List)

		return
	}

	host := fmt.Sprintf("%v:%v", server, token)

	oc.List = append(oc.List, host)

	fmt.Printf("add: %v", oc.List)

}

func (oc ocsconfig) serverExists(server string) bool {
	for i := range oc.List {
		s, _ := oc.getServerAndToken(i)
		if server == s {
			return true
		}
	}
	return false
}

func (oc ocsconfig) getServerAndToken(index int) (string, string) {
	s := strings.Split(oc.List[index], ":")

	return s[0], s[1]
}

func (oc *ocsconfig) setSelected(i int) {
	oc.Selected = i
}

func (oc *ocsconfig) updateHost(server, token string) {
	for i := range oc.List {
		s, _ := oc.getServerAndToken(i)
		if s == server {
			oc.List[i] = fmt.Sprintf("%v:%v", server, token)
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
