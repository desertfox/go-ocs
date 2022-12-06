package cmd

import (
	"fmt"
	"os"

	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "ocs"}
	config  = ocs.GetConfig()
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ocs.PrintStatus(config)
}
