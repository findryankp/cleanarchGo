package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func gitigonreCreate(featuresName string) {
	base := "./" + featuresName
	file, err := generates.FilesCreate(base + "/.gitignore")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, gitigonreContent(featuresName))
		fmt.Println("Config File Created")
	}
}

func gitigonreContent(featuresName string) string {
	var text = `local.env`
	return text
}
