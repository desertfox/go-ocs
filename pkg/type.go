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

type Ocs struct {
	Server, Token string
	config        ocsconfig
}

type ocsconfig struct {
	Selected int `yaml:"Selected"`
	List     []string
}

func (ocsc ocsconfig) getSelected() []string {
	return strings.Split(ocsc.List[ocsc.Selected], ":")
}

func (ocsc *ocsconfig) setSelected(i int) {
	ocsc.Selected = i
}

func (o Ocs) initConfig() {

	emptyOcsc := ocsconfig{0, []string{}}

	data, err := yaml.Marshal(&emptyOcsc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	err = ioutil.WriteFile(ocsconfigFile, data, 0777)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}
}

func (o *Ocs) loadConfig() {

	ocsc := &ocsconfig{}

	file, err := ioutil.ReadFile(ocsconfigFile)
	if err != nil {
		o.initConfig()
	}

	err = yaml.Unmarshal(file, &ocsc)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	o.config = *ocsc
}

func (oc *ocsconfig) add(server, token string) {

	fmt.Printf("add called: %v, %v\n", server, token)

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
