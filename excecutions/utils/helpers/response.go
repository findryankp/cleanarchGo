package helpers

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func responseCreate() {
	file, err := generates.FilesCreate("./utils/helpers/response.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, responseContent())
		fmt.Println("Cors File Created")
	}
}

func responseContent() string {
	var text = `package helpers

func ResponseSuccess(message string, data any) map[string]any {
	return map[string]any{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

func ResponseFail(message string) map[string]any {
	return map[string]any{
		"status":  false,
		"message": message,
	}
}
	
	`

	return text
}
