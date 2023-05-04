package database

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func mysqlConCreate() {
	file, err := generates.FilesCreate("./apps/database/mysqlCon.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, mysqlConContent())
		fmt.Println("Config File Created")
	}
}

func mysqlConContent() string {
	var text = `package configs
	
type AppConfig struct {
	DBUSERNAME string
	DBPASS     string
	DBHOST     string
	DBPORT     string
	DBNAME     string
	JWTKEY     string
}

func InitConfig() *AppConfig {
	return InitEnv()
}
	`

	return text
}
