package configutils

import (
	"errors"
	"github.com/spf13/viper"
)

func LoadConfigWithViper(viper *viper.Viper, configName string, dst any, configTypes ...string) error {
	viper.SetConfigName(configName)
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
	var found bool
	for _, fileType := range configTypes {
		viper.SetConfigType(fileType)
		if err := viper.ReadInConfig(); err == nil {
			found = true
			break
		}
	}
	if !found {
		return errors.New("not found config file")
	}

	if err := viper.Unmarshal(dst); err != nil {
		return err
	}
	return nil
}
