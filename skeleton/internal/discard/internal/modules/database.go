package modules

import "gorm.io/gorm"

type DatabaseModule interface {
	GetMySQL() *gorm.DB
}
