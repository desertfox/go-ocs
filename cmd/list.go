package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list hosts",
	Long:  ``,
	Run:   func(_ *cobra.Command, _ []string) {},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
