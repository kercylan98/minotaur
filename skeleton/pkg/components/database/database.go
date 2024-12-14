package database

import (
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components"
	"gorm.io/gorm"
)

var (
	_ components.DatabaseComponent     = (*databaseComponent)(nil)
	_ application.ComponentInitializer = (*databaseComponent)(nil)
)

func NewDatabaseComponent() application.Component {
	return &databaseComponent{}
}

type databaseComponent struct {
	gormDB *gorm.DB
}

func (d *databaseComponent) OnInitialize(app *application.Context) error {
	return nil
}

func (d *databaseComponent) OnImport(provider *application.ComponentProvider) {

}

func (d *databaseComponent) OnStart(app *application.Context) {

}

func (d *databaseComponent) GetMDB() *gorm.DB {
	return d.gormDB
}
