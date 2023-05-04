package middlewares

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func authenticationCreate() {
	file, err := generates.FilesCreate("./apps/middlewares/authentication.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, authenticationContent())
		fmt.Println("Authentication File Created")
	}
}

func authenticationContent() string {
	var text = `package middlewares

import (
	"` + excecutions.ModuleName + `/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ValidateToken(c); err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFail(err.Error()))
		}
		return next(c)
	}
}
	
	`

	return text
}
