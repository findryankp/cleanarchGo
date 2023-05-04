package apps

import (
	"github.com/Findryankp/cleanarchGo/excecutions/apps/configs"
	"github.com/Findryankp/cleanarchGo/excecutions/apps/database"
	"github.com/Findryankp/cleanarchGo/excecutions/apps/middlewares"
	"github.com/Findryankp/cleanarchGo/excecutions/apps/routes"
)

func InitApps() {
	configs.InitConfig()
	database.InitDatabase()
	middlewares.InitMiddlewares()
	routes.InitRoutes()
}
