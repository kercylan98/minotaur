package cmd

import (
	"github.com/spf13/cobra"
)

var tableXlsxCmd = &cobra.Command{
	Use:   "xlsx",
	Short: "Generate xlsx configuration code",
	Long:  `Generate xlsx configuration code for real-time loading and updating`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	tableCmd.AddCommand(tableXlsxCmd)
}
