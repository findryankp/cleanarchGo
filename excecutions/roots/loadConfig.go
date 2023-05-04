package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func loadConfigCreate(featuresName string) {
	base := "./" + featuresName
	file, err := generates.FilesCreate(base + "/loadConfig.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, loadConfigContent(featuresName))
		fmt.Println("loadConfig File Created")
	}
}

func loadConfigContent(featuresName string) string {
	var text = `package main

import (
	"` + generates.ModuleName + `/apps/configs"
	"` + generates.ModuleName + `/apps/database"
	"` + generates.ModuleName + `/apps/routes"

	"github.com/labstack/echo/v4"
)

func loadConfigs() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	routes.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}	
`
	return text
}
