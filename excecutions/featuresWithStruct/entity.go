package featuresWithStruct

import (
	"log"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func entityCreate(featuresName string, structString string) {
	base := "./features/" + featuresName
	file, err := generates.FilesCreate(base + "/entity.go")
	if err != nil {
		log.Println("[entityCreate] - ", err.Error())
		return
	}

	generates.FilesAddContent(file, entityContent(featuresName, structString))
	log.Println("[entityCreate] - Config File Created")
}

func entityContent(featuresName string, structString string) string {
	var text = `package ` + featuresName + `

` + structString + `

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
