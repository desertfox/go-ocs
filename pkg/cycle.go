package ocs

func (o *Ocs) cycle() {
	if len(o.Config.List) <= 1 {
		return
	}

	selected := o.Config.Selected

	if selected+1 > len(o.Config.List)-1 {
		o.Config.Selected = 0
	} else {
		o.Config.Selected++
	}

	o.Server, o.Token = o.Config.getServerAndToken(o.Config.Selected)

	o.execLogin()

}
