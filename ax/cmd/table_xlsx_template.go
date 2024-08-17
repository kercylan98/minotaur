package cmd

import (
	_ "embed"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

//go:embed table/xlsxsheet/template.xlsx
var template []byte

var (
	tableXlsxTemplateOutput string
)

var tableXlsxTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Export an xlsx template file.",
	Long:  `Exports a predefined xlsx template file, which can be used as a starting point for configuration data creation.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := filepath.Dir(tableXlsxTemplateOutput)
		checkError(os.MkdirAll(dir, os.ModePerm))

		checkError(os.WriteFile(tableXlsxTemplateOutput, template, os.ModePerm))
	},
}

func init() {
	tableXlsxCmd.AddCommand(tableXlsxTemplateCmd)

	tableXlsxTemplateCmd.Flags().StringVarP(&tableXlsxTemplateOutput, "output", "o", "", "output file path")

	checkError(tableXlsxTemplateCmd.MarkFlagRequired("output"))
}
