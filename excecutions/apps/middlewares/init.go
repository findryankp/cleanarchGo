package middlewares

import "github.com/Findryankp/cleanarchGo/excecutions/generates"

func InitMiddlewares() {
	packagesInstall := []string{
		"github.com/labstack/echo/v4",
		"github.com/labstack/echo/v4/middleware",
		"github.com/golang-jwt/jwt/v5",
		"github.com/labstack/echo-jwt/v4",
	}
	generates.PackageIntallList(packagesInstall)

	corsCreate()
	basicLogCreate()
	authenticationCreate()
	jwtCreate()
}
