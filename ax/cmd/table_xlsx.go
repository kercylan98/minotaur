package cmd

import (
	"github.com/spf13/cobra"
)

var tableXlsxCmd = &cobra.Command{
	Use:   "xlsx",
	Short: "Generates configuration code from xlsx files.",
	Long:  `Generates configuration code from xlsx files, enabling real-time loading and updates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	tableCmd.AddCommand(tableXlsxCmd)
}
