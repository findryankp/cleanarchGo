package data

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func queryCreate() {
	name := "./features/users/data/query.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, queryContent())
		fmt.Println("Query File Created")
	}
}

func queryContent() string {
	var text = `package data

import (
	"` + generates.ModuleName + `/features/users` + `"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]users.Core, error) {
	var dataModel []Users
	if err := q.db.Find(&dataModel); err.Error != nil {
		return nil, err.Error
	}
	return ListModelToCore(dataModel), nil
}

func (q *query) SelectById(id uint) (users.Core, error) {
	var dataModel Users
	if err := q.db.First(&dataModel, id); err.Error != nil {
		return users.Core{}, err.Error
	}
	return ModelToCore(dataModel), nil
}

func (q *query) SelectByEmail(email string) (users.Core, error) {
	var dataModel Users
	if err := q.db.Where("email = ?", email).First(&dataModel); err.Error != nil {
		return users.Core{}, err.Error
	}
	return ModelToCore(dataModel), nil
}

func (q *query) Store(dataCore users.Core) (uint, error) {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Create(&dataModel); err.Error != nil {
		return 0, err.Error
	}
	return dataModel.ID, nil
}

func (q *query) Edit(dataCore users.Core, id uint) error {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Where("id", id).Updates(&dataModel); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var dataModel Users
	if err := q.db.Delete(&dataModel, id); err.Error != nil {
		return err.Error
	}
	return nil
}
	`

	return text
}
