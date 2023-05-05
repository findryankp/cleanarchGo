package additions

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func routeCreate(featuresName string) {
	name := "./apps/routes/" + featuresName + "Router.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, routeContent(featuresName))
		fmt.Println("Request File Created")
	}

	lineNumber := generates.ContentGetLinenumber("./apps/routes/routes.go", `middlewares.BasicLogger(e)`)

	generates.ContentInsertAtLinenumber("./apps/routes/routes.go", `	middlewares.BasicLogger(e)`+"\n"+`	`+featuresName+"Router(db,e)", lineNumber-1)

	fmt.Println("Routes successfully created")
}

func routeContent(featuresName string) string {
	var text = `package routes

import (
	` + featuresName + `Data "` + generates.ModuleName + `/features/` + featuresName + `/data"
	` + featuresName + `Handler "` + generates.ModuleName + `/features/` + featuresName + `/delivery"
	` + featuresName + `Service "` + generates.ModuleName + `/features/` + featuresName + `/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ` + featuresName + `Router(db *gorm.DB, e *echo.Echo) {
	data := ` + featuresName + `Data.New(db)
	service := ` + featuresName + `Service.New(data)
	handler := ` + featuresName + `Handler.New(service)

	g := e.Group("/` + generates.ToLower(featuresName) + `")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
`
	return text
}
