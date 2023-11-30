package featuresWithStruct

import (
	"github.com/Findryankp/cleanarchGo/excecutions/featuresWithStruct/additions"
	"github.com/Findryankp/cleanarchGo/excecutions/featuresWithStruct/data"
	"github.com/Findryankp/cleanarchGo/excecutions/featuresWithStruct/delivery"
	"github.com/Findryankp/cleanarchGo/excecutions/featuresWithStruct/service"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitFeatures(featureName string) {
	createFolderFeature(featureName)

	//create entity
	entityCreate(featureName, "")

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
