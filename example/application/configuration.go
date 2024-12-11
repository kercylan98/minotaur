package application

import "github.com/gofiber/fiber/v2"

func newConfiguration() *Configuration {
	return &Configuration{}
}

type Configuration struct {
	addr                string
	fiberSettingHandler func(fiberApp *fiber.App)
}

func (c *Configuration) WithAddr(addr string) *Configuration {
	c.addr = addr
	return c
}

func (c *Configuration) WithFiberSettings(handler func(fiberApp *fiber.App)) *Configuration {
	c.fiberSettingHandler = handler
	return c
}
