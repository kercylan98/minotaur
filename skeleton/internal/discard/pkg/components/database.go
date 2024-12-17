package components

import "gorm.io/gorm"

type DatabaseComponent interface {
	GetMDB() *gorm.DB
}
