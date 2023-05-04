package middlewares

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func jwtCreate() {
	file, err := generates.FilesCreate("./apps/middlewares/jwt.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, jwtContent())
		fmt.Println("JWT File Created")
	}
}

func jwtContent() string {
	var text = ``

	return text
}
