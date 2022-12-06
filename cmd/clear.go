package cmd

import (
	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear saved hosts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ocs.WriteConfig(ocs.EmptyConfig())
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
