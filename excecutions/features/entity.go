package features

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func entityCreate(featuresName string) {
	base := "./features/" + featuresName
	file, err := generates.FilesCreate(base + "/entity.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, entityContent(featuresName))
		fmt.Println("Config File Created")
	}
}

func entityContent(featuresName string) string {
	titleCase := strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, featuresName)

	var text = `package ` + featuresName + `

import "time"

type Core struct {
	Id        uint
	` + titleCase + `  string ` + "`validate:" + `"required"` + "`" + `
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) error
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	Store(data Core) (uint, error)
	Edit(data Core, id uint) error
	Destroy(id uint) error
}
	`

	return text
}
