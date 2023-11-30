package additions

import (
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func migrationAdd(featuresName string) {
	file := "./apps/database/migration.go"
	line := generates.ContentGetLinenumber(file, "gorm.io/gorm")

	//add import file
	generates.ContentInsertAtLinenumber(file, migrationContentImport(featuresName), line-1)

	//add content migration
	line = generates.ContentGetLinenumber(file, "InitMigration(db *gorm.DB) ")
	generates.ContentInsertAtLinenumber(file, migrationContent(featuresName), line-1)
}

func migrationContentImport(featuresName string) string {
	var text = `	"gorm.io/gorm"
	_` + featuresName + `"` + generates.ModuleName + `/features/` + featuresName + `/data"
	`
	return text
}

func migrationContent(featuresName string) string {
	TitleCase := generates.ToTitle(featuresName)
	var text = `func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&_` + featuresName + `.` + TitleCase + `{})`
	return text
}
