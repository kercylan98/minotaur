package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minotaur-ax",
	Short: "Minotaur ax is a CLI (Command Line Interface) tool designed to facilitate the rapid development of Minotaur projects.",
	Long:  "Minotaur ax is an essential command line tool aimed at enhancing and accelerating the development process of projects using the Minotaur framework.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {

}
