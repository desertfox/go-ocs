package ocs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Ocs struct {
	Server, Token string
	config        ocsconfig
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
