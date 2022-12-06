package cmd

import (
	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "prune more than 24hr old hosts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.Prune()

		if len(config.Hosts) > 0 {
			if err := ocs.Login(config); err != nil {
				config.Del(config.Selected)
				ocs.Login(config)
			}
		}

		ocs.WriteConfig(config)
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}
