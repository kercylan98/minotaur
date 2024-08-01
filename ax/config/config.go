package config

import (
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/ax/constants"
	"github.com/spf13/viper"
)

const (
	Type = "toml"
	Name = ".minotaur"
)

func init() {
	viper.AddConfigPath(constants.UserHomeDir())
	viper.SetConfigType(Type)
	viper.SetConfigName(Name)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Using default config")
	}
}
