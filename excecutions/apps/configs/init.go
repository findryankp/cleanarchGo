package configs

import "github.com/Findryankp/cleanarchGo/excecutions/generates"

func InitConfig() {
	packagesInstall := []string{
		"github.com/spf13/viper",
		"github.com/go-playground/validator/v10",
	}
	generates.PackageIntallList(packagesInstall)

	//generate file
	configCreate()
	configEnvCreate()
}
