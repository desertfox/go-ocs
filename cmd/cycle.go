package cmd

import (
	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var cycleCmd = &cobra.Command{
	Use:   "",
	Short: "cycle logged in host",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.CycleSelected()

		if err := ocs.Login(config); err != nil {
			config.Del(config.Selected)
			ocs.Login(config)
		}

		ocs.WriteConfig(config)
	},
}

func init() {
	rootCmd.AddCommand(cycleCmd)
}
