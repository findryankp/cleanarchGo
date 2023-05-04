package data

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func modelCreate(featuresName string) {
	name := "./features/" + featuresName + "/data/model.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, modelContent(featuresName))
		fmt.Println("Config File Created")
	}
}

func modelContent(featuresName string) string {
	TitleCase := strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, featuresName)

	var text = `package data

import (
	"` + generates.ModuleName + `/features/` + featuresName + `"

	"gorm.io/gorm"
)

type ` + TitleCase + ` struct {
	gorm.Model
	` + TitleCase + ` string
}

func CoreToModel(coreData ` + featuresName + `.Core) ` + TitleCase + ` {
	return ` + TitleCase + `{
		` + TitleCase + `: coreData.` + TitleCase + `,
	}
}

func ModelToCore(modelData ` + TitleCase + `) ` + featuresName + `.Core {
	return ` + featuresName + `.Core{
		Id:        modelData.ID,
		` + TitleCase + `:  modelData.` + TitleCase + `,
		CreatedAt: modelData.CreatedAt,
		UpdatedAt: modelData.UpdatedAt,
	}
}

func ListModelToCore(modelData []` + TitleCase + `) []` + featuresName + `.Core {
	var coreData []` + featuresName + `.Core
	for _, v := range modelData {
		coreData = append(coreData, ModelToCore(v))
	}
	return coreData
}
	`

	return text
}
