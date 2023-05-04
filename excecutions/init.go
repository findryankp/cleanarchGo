package excecutions

import (
	"github.com/Findryankp/cleanarchGo/excecutions/apps"
	"github.com/Findryankp/cleanarchGo/excecutions/generates"
	"github.com/Findryankp/cleanarchGo/excecutions/utils"
)

func InitExecution() {
	generates.InitModule()
	apps.InitApps()
	utils.InitUtils()
}
