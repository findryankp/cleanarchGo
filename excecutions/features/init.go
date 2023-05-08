package features

import (
	"github.com/Findryankp/cleanarchGo/excecutions/features/additions"
	"github.com/Findryankp/cleanarchGo/excecutions/features/data"
	"github.com/Findryankp/cleanarchGo/excecutions/features/delivery"
	"github.com/Findryankp/cleanarchGo/excecutions/features/service"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitFeatures(featureName string) {
	createFolderFeature(featureName)
	entityCreate(featureName)

	data.InitData(featureName)
	delivery.InitDelivery(featureName)
	service.InitService(featureName)
	additions.InitAdditions(featureName)
}

func createFolderFeature(featuresName string) {
	base := "./features/" + featuresName
	folderName := []string{
		base,
		base + "/data",
		base + "/delivery",
		base + "/service",
	}
	generates.FolderCreateList(folderName)
}
