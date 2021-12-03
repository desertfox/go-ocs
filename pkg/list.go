package ocs

import "fmt"

func (o Ocs) list() {
	for i, v := range o.Config.List {
		fmt.Printf("%v:%v\n", i, v)
	}
}
