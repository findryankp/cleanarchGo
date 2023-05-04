package database

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func mysqlConCreate() {
	file, err := generates.FilesCreate("./apps/database/mysqlCon.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, mysqlConContent())
		fmt.Println("mysqlCon File Created")
	}
}

func mysqlConContent() string {
	var text = `package database

import (
	"` + excecutions.ModuleName + `/apps/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(app configs.AppConfig) *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", app.DBUSERNAME, app.DBPASS, app.DBHOST, app.DBPORT, app.DBNAME)

	DB, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return DB
}
	
	`

	return text
}
