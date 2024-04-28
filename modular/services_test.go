package modular_test

import (
	"github.com/kercylan98/minotaur/modular"
)

type AccountService struct {
	config ConfigServiceExposer
}

func (a *AccountService) OnInit(application *modular.Application) {
}

func (a *AccountService) OnPreload(application *modular.Application) {
	a.config = modular.InvokeService[ConfigServiceExposer](application)
}

func (a *AccountService) OnMount(application *modular.Application) {
}

func (a *AccountService) OnStart(application *modular.Application) {
	// 假设需要使用配置
	a.config.Get("key")
}

func (a *AccountService) Login() {

}

type AccountServiceExposer interface {
	Login()
}

type ConfigService struct {
	kv map[string]string
}

func (c *ConfigService) OnInit(application *modular.Application) {
}

func (c *ConfigService) OnPreload(application *modular.Application) {
}

func (c *ConfigService) OnMount(application *modular.Application) {
	// 假设从数据库加载
	c.kv = make(map[string]string)
}

func (c *ConfigService) OnStart(application *modular.Application) {
}

func (c *ConfigService) Get(key string) string {
	if c.kv == nil {
		panic("config service not initialized")
	}
	return key
}

type ConfigServiceExposer interface {
	Get(key string) string
}
