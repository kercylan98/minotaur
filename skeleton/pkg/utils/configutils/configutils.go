package configutils

import (
	"fmt"
	"strings"
)

func CheckLoggerLevel(level string) error {
	loggerLevel := strings.ToLower(level)
	switch loggerLevel {
	case "debug", "info", "warn", "error":
	default:
		return fmt.Errorf("invalid logger level: %s, only support debug, info, warn, error", loggerLevel)
	}
	return nil
}

func CheckEnv(env string) error {
	env = strings.ToLower(env)
	switch env {
	case "dev", "prod", "test":
	default:
		return fmt.Errorf("invalid environment: %s, only support dev, prod, test", env)
	}
	return nil
}
