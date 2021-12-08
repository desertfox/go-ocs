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

const version = "0.0.7"

func init() {
	if os.Getenv("OCS_DEBUG") == "1" {
		selfupdate.EnableLog()
	}

	doSelfUpdate()

	if len(os.Args) > 1 {
		CLICommand = os.Args[1]
	} else {
		CLICommand = ""
	}

	FlagSet = flag.FlagSet{}

	FlagSet.StringVar(&CLIServer, "server", "", "Server")
	FlagSet.StringVar(&CLIToken, "token", "", "Token")

}

func main() {

	if len(os.Args) > 1 {
		FlagSet.Parse(os.Args[2:])

		if os.Args[1] == "help" || os.Args[1] == "h" {
			fmt.Println(readme)
			os.Exit(0)
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

	ocs.DoCommand(CLICommand)
}

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "desertfox/go-ocs")
	if err != nil {
		log.Println("Unable to execute update: ", err)
		return
	}

	if latest.Version.Equals(v) {
		//NO-OP
	} else {
		log.Println("Successfully updated ocs to version: ", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
		log.Println("Re-run your command.")
		os.Exit(0)
	}
}
