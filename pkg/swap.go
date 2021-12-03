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

	if i > len(o.Config.List)-1 {
		fmt.Printf("Swap %v greater than %v of config values", i, len(o.Config.List)-1)
		return
	} else {
		o.Config.setSelected(i)
	}

	o.Server, o.Token = o.Config.getServerAndToken(o.Config.Selected)

	o.execLogin()

}
