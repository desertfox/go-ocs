package ocs

func (o Ocs) DoCommand(CLICommand string) {

	o.loadConfig()

	switch CLICommand {

	case "login":
		o.add()
	case "swap":
		o.swap()
	case "list":
		o.list()
	default:
		o.cycle()
	}
}
