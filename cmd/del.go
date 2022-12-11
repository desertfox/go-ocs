package cmd

import (
	"strconv"

	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete logged in host",
	Long:  ``,
	Run: func(_ *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])

		config.Del(i)

		ocs.WriteConfig(config)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
