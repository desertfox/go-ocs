package cmd

import (
	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear saved hosts",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		ocs.WriteConfig(ocs.EmptyConfig())
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
