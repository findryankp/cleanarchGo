package middlewares

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func basicLogCreate() {
	file, err := generates.FilesCreate("./apps/middlewares/basicLog.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, basicLogContent())
		fmt.Println("basicLog File Created")
	}
}

func basicLogContent() string {
	var text = `package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicLogger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
	
	`

	return text
}
