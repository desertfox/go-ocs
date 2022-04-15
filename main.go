package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/blang/semver"
	ocs "github.com/desertfox/ocs/pkg"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var (
	CLICommand, CLIToken, CLIServer string
	FlagSet                         flag.FlagSet
)

var (
	//go:embed README.md
	readme string
)

const version = "0.1.2"

func init() {
	if os.Getenv("OCS_DEBUG") == "1" {
		selfupdate.EnableLog()
	}

	FlagSet = flag.FlagSet{}
	FlagSet.StringVar(&CLIServer, "server", "", "Server")
	FlagSet.StringVar(&CLIToken, "token", "", "Token")
}

func main() {
	var opts []string = os.Args

	if len(opts) > 1 {
		if opts[0] == "oc" {
			opts = opts[1:]
		}

		CLICommand = opts[1]

		FlagSet.Parse(opts[2:])

		if CLICommand == "help" || CLICommand == "h" {
			fmt.Println(readme)
			os.Exit(0)
		}

		if CLICommand == "login" {
			CLIServer = opts[2]
			FlagSet.Parse(opts[3:])
		}
	}

	host := ocs.Host{
		Server:  CLIServer,
		Token:   CLIToken,
		Created: time.Now(),
	}
	config := ocs.GetConfig()

	ocs := ocs.Ocs{
		Host:   host,
		Config: config,
	}

	doSelfUpdate(ocs)

	ocs.DoCommand(CLICommand)
}

func doSelfUpdate(ocs ocs.Ocs) {
	waitPeriodMin := time.Now().Add(-1 * 24 * time.Hour)

	if waitPeriodMin.After(ocs.Config.UpdateCheck) {
		v := semver.MustParse(version)
		latest, err := selfupdate.UpdateSelf(v, "desertfox/go-ocs")
		if err != nil {
			log.Println("Unable to execute update: ", err)
			return
		}

		ocs.SetUpdateCheck()

		if latest.Version.Equals(v) {
			//NO-OP
		} else {
			log.Println("Successfully updated ocs to version: ", latest.Version)
			log.Println("Release note:\n", latest.ReleaseNotes)
			log.Println("Re-run your command.")
			os.Exit(0)
		}
	}

}
