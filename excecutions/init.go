package excecutions

import "github.com/Findryankp/cleanarchGo/excecutions/generates"

var ModuleName string

func InitExcecution() {
	moduleName, _ := generates.ModuleNameGet()
	ModuleName = moduleName
}
