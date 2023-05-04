package delivery

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func requestCreate(featuresName string) {
	name := "./features/" + featuresName + "/delivery/request.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, requestContent(featuresName))
		fmt.Println("Request File Created")
	}
}

func requestContent(featuresName string) string {
	TitleCase := strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, featuresName)

	var text = `package delivery

import "` + generates.ModuleName + `/features/` + featuresName + `"

type Request struct {
	` + TitleCase + ` string ` + "`json:" + featuresName + `"` + "` " + "`form:" + featuresName + `"` + "`" + `
}

func RequestToCore(request *Request) ` + featuresName + `.Core {
	return ` + featuresName + `.Core{
		` + TitleCase + `: request.` + TitleCase + `,
	}
}
	
`
	return text
}
