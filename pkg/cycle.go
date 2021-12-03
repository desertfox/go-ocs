package ocs

func (o *Ocs) cycle() {
	if len(o.Config.Hosts) <= 1 {
		return
	}

	selected := o.Config.Selected

	if selected+1 > len(o.Config.Hosts)-1 {
		o.Config.Selected = 0
	} else {
		o.Config.Selected++
	}

	o.Host = o.Config.GetSelectedHost()

	o.execLogin()

}
