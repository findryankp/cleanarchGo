package users

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func entityCreate() {
	base := "./features/users"
	file, err := generates.FilesCreate(base + "/entity.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, entityContent())
		fmt.Println("Config File Created")
	}
}

func entityContent() string {
	var text = `package users

import "time"

type Core struct {
	Id        uint
	Email     string ` + "`" + `validate:"required"` + "`" + `
	Password  string ` + "`" + `validate:"required"` + "`" + `
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) error

	Login(email, password string) (Core, error)
	ChangePassword(id uint, old, new, confirm, hash string) error
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	SelectByEmail(email string) (Core, error)
	Store(data Core) (uint, error)
	Edit(data Core, id uint) error
	Destroy(id uint) error
}
	`

	return text
}
