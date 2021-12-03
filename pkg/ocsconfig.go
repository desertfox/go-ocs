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

func (oc ocsconfig) getSelected() []string {
	return strings.Split(oc.List[oc.Selected], ":")[0:1]
}

func (oc *ocsconfig) setSelected(i int) {
	oc.Selected = i
}

func (oc *ocsconfig) add(server, token string) {

	oc.List = append(oc.List, fmt.Sprintf("%v:%v", server, token))

	fmt.Printf("add: %v", oc.List)

	oc.writeConfig()
}

func (oc ocsconfig) writeConfig() {

	ocsc := oc

	data, err := yaml.Marshal(&ocsc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	err = ioutil.WriteFile(filepath.Join(wd, ocsconfigFile), data, 0777)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}
}
