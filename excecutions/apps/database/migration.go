package database

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func migrationCreate() {
	file, err := generates.FilesCreate("./apps/database/migration.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, migrationContent())
		fmt.Println("Migration File Created")
	}
}

func migrationContent() string {
	var text = `package database

import (
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	fmt.Println("Migration Done")
}
	
	`

	return text
}
