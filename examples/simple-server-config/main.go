// 该案例演示了配置导表工具的使用，其中包括了作为模板的配置文件及导出的配置文件
package main

import (
	"github.com/kercylan98/minotaur/config"
	"github.com/kercylan98/minotaur/examples/simple-server-config/config/configs"
	"github.com/kercylan98/minotaur/planner/configexport"
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
	"path/filepath"
)

const (
	workdir = "./examples/simple-server-config"
)

func main() {
	export()
	config.Init(filepath.Join(workdir, "config", "json"), configs.LoadConfig, configs.Refresh)
	config.Load()
	config.Refresh()
	log.Info("Config", zap.Any("SystemConfig", configs.ISystemConfig))
	log.Info("Config", zap.Any("WelcomeConfig", configs.IWelcomeConfig))
}

func export() {
	c := configexport.New(filepath.Join(workdir, "config", "系统配置.xlsx"))
	c.ExportGo("", filepath.Join(workdir, "config", "configs"))
	c.ExportServer("", filepath.Join(workdir, "config", "json"))
}
