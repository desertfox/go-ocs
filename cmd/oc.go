package cmd

import (
	"github.com/desertfox/ocs/pkg/ocs"
	"github.com/spf13/cobra"
)

var (
	ocCmd = &cobra.Command{
		Use:   "oc",
		Short: "oc login command string",
		Run: func(cmd *cobra.Command, args []string) {
			h := ocs.NewHost(server, token)

			config.Add(h)

			ocs.Login(config)

			ocs.WriteConfig(config)
		},
	}
	token, server string
)

func init() {
	rootCmd.AddCommand(ocCmd)
	ocCmd.PersistentFlags().StringVar(&server, "server", "", "Server")
	ocCmd.PersistentFlags().StringVar(&token, "token", "", "Token")
}
