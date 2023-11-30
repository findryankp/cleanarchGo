package data

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func queryCreate(featuresName string) {
	name := "./features/" + featuresName + "/data/query.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, queryContent(featuresName))
		fmt.Println("Query File Created")
	}
}

func queryContent(featuresName string) string {
	TitleCase := generates.ToTitle(featuresName)

	var text = `package data

import (
	"` + generates.ModuleName + `/features/` + featuresName + `"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) ` + featuresName + `.DataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]` + featuresName + `.Core, error) {
	var dataModel []` + TitleCase + `
	if err := q.db.Find(&dataModel); err.Error != nil {
		return nil, err.Error
	}
	return ListModelToCore(dataModel), nil
}

func (q *query) SelectById(id uint) (` + featuresName + `.Core, error) {
	var dataModel ` + TitleCase + `
	if err := q.db.First(&dataModel, id); err.Error != nil {
		return ` + featuresName + `.Core{}, err.Error
	}
	return ModelToCore(dataModel), nil
}

func (q *query) Store(dataCore ` + featuresName + `.Core) (uint, error) {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Create(&dataModel); err.Error != nil {
		return 0, err.Error
	}
	return dataModel.ID, nil
}

func (q *query) Edit(dataCore ` + featuresName + `.Core, id uint) error {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Where("id", id).Updates(&dataModel); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var dataModel ` + TitleCase + `
	if err := q.db.Delete(&dataModel, id); err.Error != nil {
		return err.Error
	}
	return nil
}
	`

	return text
}
