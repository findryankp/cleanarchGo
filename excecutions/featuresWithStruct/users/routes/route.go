package routes

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func UserRoutesCreate() {
	generates.FilesDelete("./apps/routes/usersRouter.go")
	file, err := generates.FilesCreate("./apps/routes/usersRouter.go")
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
	"` + generates.ModuleName + `/apps/middlewares"
	usersData "` + generates.ModuleName + `/features/users/data"
	usersHandler "` + generates.ModuleName + `/features/users/delivery"
	usersService "` + generates.ModuleName + `/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func usersRouter(db *gorm.DB, e *echo.Echo) {
	data := usersData.New(db)
	service := usersService.New(data)
	handler := usersHandler.New(service)

	e.POST("login", handler.Login)
	e.POST("register", handler.Register)
	e.POST("change-password", handler.ChangePassword, middlewares.Authentication)
	e.GET("claim-token", handler.ClaimsToken, middlewares.Authentication)

	g := e.Group("/users")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.PUT("/:id", handler.Update)
}

	`

	return text
}
