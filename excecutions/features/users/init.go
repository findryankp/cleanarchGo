package users

import (
	"github.com/Findryankp/cleanarchGo/excecutions/apps/routes"
	"github.com/Findryankp/cleanarchGo/excecutions/features/additions"
	"github.com/Findryankp/cleanarchGo/excecutions/features/users/data"
	"github.com/Findryankp/cleanarchGo/excecutions/features/users/delivery"
	"github.com/Findryankp/cleanarchGo/excecutions/features/users/service"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitUsers() {
	packagesInstall := []string{
		"github.com/go-playground/validator/v10",
	}
	generates.PackageIntallList(packagesInstall)

	createFolderFeature()
	entityCreate()

	data.InitData()
	delivery.InitDelivery()
	service.InitService()
	routes.InitRoutes()
	additions.InitAdditions("users")
}

func createFolderFeature() {
	base := "./features/users"
	folderName := []string{
		base,
		base + "/data",
		base + "/delivery",
		base + "/service",
	}
	generates.FolderCreateList(folderName)
}
