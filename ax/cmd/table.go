package cmd

import (
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Generates code and exports data from configuration tables.",
	Long:  `Generates code and exports data from configuration tables with support for various export formats.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
