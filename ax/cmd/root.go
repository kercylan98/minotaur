package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var stack bool

var rootCmd = &cobra.Command{
	Use:   "minotaur-ax",
	Short: "Minotaur ax is a CLI (Command Line Interface) tool designed to facilitate the rapid development of Minotaur projects.",
	Long:  "Minotaur ax is an essential command line tool aimed at enhancing and accelerating the development process of projects using the Minotaur framework.",
}

func Execute() {
	defer func() {
		if err := recover(); err != nil {
			checkError(err)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&stack, "stack", false, "show stack trace")
}

func checkError(err any) {
	if err != nil {
		if stack {
			panic(err)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
