package ocs

func (o *Ocs) cycle() {
	selected := o.config.Selected

	if selected+1 > len(o.config.List)-1 {
		o.config.Selected = 0
	} else {
		o.config.Selected++
	}

	data := o.config.getSelected()

	o.Server = data[0]
	o.Token = data[1]

	o.execLogin()

	o.config.writeConfig()

}
