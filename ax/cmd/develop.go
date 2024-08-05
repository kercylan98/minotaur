package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	configKeyDevelopWorkdir = "develop.workdir"
)

// developCmd represents the develop command
var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "CLI commands involved in minotaur development",
	Long:  "Use this command to develop minotaur, such as building, testing, etc.",
	RunE: func(cmd *cobra.Command, args []string) error {
		abs, err := filepath.Abs(viper.GetString(configKeyDevelopWorkdir))
		cobra.CheckErr(err)

		fmt.Println("minotaur develop workdir:", abs)
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(developCmd)

	viper.SetDefault(configKeyDevelopWorkdir, "../")
}
