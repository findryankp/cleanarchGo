package excecutions

import (
	"github.com/Findryankp/cleanarchGo/excecutions/apps"
	"github.com/Findryankp/cleanarchGo/excecutions/features"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
	"github.com/Findryankp/cleanarchGo/excecutions/roots"
	"github.com/Findryankp/cleanarchGo/excecutions/utils"
)

func CommandInit() {
	generates.InitModule()
	generates.FolderInit()
	apps.InitApps()
	utils.InitUtils()
	roots.InitRoots()
}

func NewFeatures(featuresName string) {
	generates.InitModule()
	features.InitFeatures(featuresName)
}
