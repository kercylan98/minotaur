package cmd

import (
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Used for data configuration table code generation and data export.",
	Long:  `Used for data configuration table code generation and data export. It supports multiple types of exports`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
