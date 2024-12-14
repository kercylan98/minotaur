package database

import (
	"github.com/kercylan98/minotaur/skeleton/internal/modules"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_ module.Dependency[*application.Context]      = (*databaseModule[*application.Context])(nil)
	_ module.DependencySetup[*application.Context] = (*databaseModule[*application.Context])(nil)
)

func NewDatabaseModule() module.Module[*application.Context] {
	return &databaseModule{}
}

type databaseModule struct {
	app      *application.Context
	services modules.DatabaseModule
	mysql    *gorm.DB

	configuratorModule modules.ConfiguratorModule
}

func (d *databaseModule) OnInitialize(app *application.Context) error {
	d.app = app
	d.services = newDatabaseServices(d)
	return nil
}

func (d *databaseModule) OnDependencyInitialize(loader *module.Loader[*application.Context]) error {
	d.configuratorModule = module.Load[*application.Context, modules.ConfiguratorModule](loader)
	return nil
}

func (d *databaseModule) OnDependencySetup() error {
	if err := d.openMySQLConnection(); err != nil {
		return err
	}
	return nil
}

func (d *databaseModule) openMySQLConnection() error {
	mysqlConfig := d.configuratorModule.GetRuntimeConfig().Database.MySQL
	if !mysqlConfig.Enabled {
		return nil
	}
	dsn := mysqlConfig.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.mysql = db
	return nil
}
