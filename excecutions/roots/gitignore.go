package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func gitigonreCreate() {
	file, err := generates.FilesCreate("./.gitignore")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, gitigonreContent())
		fmt.Println("Gitignore File Created")
	}
}

func gitigonreContent() string {
	var text = `local.env`
	return text
}
