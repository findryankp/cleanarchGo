package data

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func modelCreate() {
	name := "./features/users/data/model.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, modelContent())
		fmt.Println("Model File Created")
	}
}

func modelContent() string {
	var text = `package data

import (
	"` + generates.ModuleName + `/features/users` + `"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string ` + "`" + `gorm:"unique"` + "`" + `
	Password string
	Name     string
	Role     string
}

func CoreToModel(coreData users.Core) Users {
	return Users{
		Email:    coreData.Email,
		Password: coreData.Password,
		Name:     coreData.Name,
		Role:     coreData.Role,
	}
}

func ModelToCore(modelData Users) users.Core {
	return users.Core{
		Id:        modelData.ID,
		Email:     modelData.Email,
		Password:  modelData.Password,
		Name:      modelData.Name,
		Role:      modelData.Role,
		CreatedAt: modelData.CreatedAt,
		UpdatedAt: modelData.UpdatedAt,
	}
}

func ListModelToCore(modelData []Users) []users.Core {
	var coreData []users.Core
	for _, v := range modelData {
		coreData = append(coreData, ModelToCore(v))
	}
	return coreData
}
	`

	return text
}
