package ocs

import (
	"fmt"
	"os"
	"strconv"
)

func (o *Ocs) swap() {

	id := os.Args[2]

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Uh oh:" + err.Error())
	}

	if i > len(o.config.List)-1 {
		fmt.Printf("Swap %v greater than %v of config values", i, len(o.config.List)-1)
		return
	} else {
		o.config.setSelected(i)
	}

	data := o.config.getSelected()

	o.Server = data[0]
	o.Token = data[1]

	o.execLogin()

}
