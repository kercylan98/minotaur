package constants

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	userHomeDir   string
	configFileDir string
)

func init() {
	uh, err := os.UserHomeDir()
	cobra.CheckErr(err)

	userHomeDir = uh
}

func UserHomeDir() string {
	return userHomeDir
}
