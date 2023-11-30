package delivery

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func requestCreate() {
	name := "./features/users/delivery/request.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, requestContent())
		fmt.Println("Request File Created")
	}
}

func requestContent() string {

	var text = `package delivery

import "` + generates.ModuleName + `/features/users` + `"

type LoginRequest struct {
	Email    string ` + "`" + `json:"email" form:"email"` + "`" + `
	Password string ` + "`" + `json:"password" form:"password"` + "`" + `
}

type Register struct {
	Email    string ` + "`" + `json:"email" form:"email"` + "`" + `
	Password string ` + "`" + `json:"password" form:"password"` + "`" + `
	Name     string ` + "`" + `json:"name" form:"name"` + "`" + `
}

type EditPassword struct {
	OldPassword     string ` + "`" + `json:"old_password" form:"old_password"` + "`" + `
	NewPassword     string ` + "`" + `json:"new_password" form:"new_password"` + "`" + `
	ConfirmPassword string ` + "`" + `json:"confirm_password" form:"confirm_password"` + "`" + `
}

type Request struct {
	Email    string ` + "`" + `json:"email" form:"email"` + "`" + `
	Password string ` + "`" + `json:"password" form:"password"` + "`" + `
	Name     string ` + "`" + `json:"name" form:"name"` + "`" + `
	Role     string ` + "`" + `json:"role" form:"role"` + "`" + `
}

func RequestToCore(request *Request) users.Core {
	return users.Core{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
}
	
`
	return text
}
