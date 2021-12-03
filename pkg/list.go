package ocs

import "fmt"

func (o Ocs) list() {
	for i, v := range o.Config.Hosts {
		fmt.Printf("%v:%#v\n", i, v)
	}
}
