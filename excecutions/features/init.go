package features

import (
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitFeatures(featureName string) {
	createFolderFeature(featureName)
	entityCreate(featureName)
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
