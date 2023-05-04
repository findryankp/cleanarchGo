package delivery

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func responseCreate(featuresName string) {
	name := "./features/" + featuresName + "/delivery/response.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, responseContent(featuresName))
		fmt.Println("Handler File Created")
	}
}

func responseContent(featuresName string) string {
	TitleCase := generates.ToTitle(featuresName)

	var text = `package delivery

import "` + generates.ModuleName + `/features/` + featuresName + `"

type Response struct {
	Id       uint   ` + "`json:" + `"id"` + "`" + `
	` + TitleCase + ` string ` + "`json:" + `"` + featuresName + `"` + "`" + `
}

func CoreToResponse(dataCore ` + featuresName + `.Core) Response {
	return Response{
		Id:       dataCore.Id,
		` + TitleCase + `: dataCore.` + TitleCase + `,
	}
}

func ListCoreToResponse(dataCore []` + featuresName + `.Core) []Response {
	var dataResponse []Response
	for _, v := range dataCore {
		dataResponse = append(dataResponse, CoreToResponse(v))
	}
	return dataResponse
}
	
`
	return text
}
