package features

import (
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func InitFeatures(featureName string) {
	CreateFolderFeature(featureName)
	entityCreate(featureName)
}

func CreateFolderFeature(featuresName string) {
	base := "./features/" + featuresName
	folderName := []string{
		base,
		base + "/data",
		base + "/delivery",
		base + "/service",
	}
	generates.FolderCreateList(folderName)
}
