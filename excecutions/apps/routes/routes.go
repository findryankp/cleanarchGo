package routes

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func routesCreate() {
	file, err := generates.FilesCreate("./apps/routes/routes.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, routesContent())
		fmt.Println("Cors File Created")
	}
}

func routesContent() string {
	var text = `package routes

import (
	"` + excecutions.ModuleName + `/apps/middlewares"
	"` + excecutions.ModuleName + `/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	e.GET("/", index)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to Cleanarch Starter Pack",
		"developmen_by": "Findryankp",
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}
	`

	return text
}
