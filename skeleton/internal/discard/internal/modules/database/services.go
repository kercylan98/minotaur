package database

import (
	"github.com/kercylan98/minotaur/skeleton/internal/discard/internal/modules"
	"gorm.io/gorm"
)

func newDatabaseServices(module *databaseModule) modules.DatabaseModule {
	return &databaseServices{
		module: module,
	}
}

type databaseServices struct {
	module *databaseModule
}

func (d *databaseServices) GetMySQL() *gorm.DB {
	return d.module.mysql
}
