package database

import (
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitDatabase() {
	packagesInstall := []string{
		"gorm.io/gorm",
		"gorm.io/driver/mysql",
	}
	generates.PackageIntallList(packagesInstall)

	mysqlConCreate()
	migrationCreate()
}
