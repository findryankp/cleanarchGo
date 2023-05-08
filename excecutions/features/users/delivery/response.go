package delivery

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func responseCreate() {
	name := "./features/users/delivery/response.go"
	file, err := generates.FilesCreate(name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, responseContent())
		fmt.Println("Handler File Created")
	}
}

func responseContent() string {

	var text = `package delivery

import "` + generates.ModuleName + `/features/users` + `"

type Response struct {
	Id    uint   ` + "`" + `json:"id"` + "`" + `
	Email string ` + "`" + `json:"email"` + "`" + `
	Name  string ` + "`" + `json:"name"` + "`" + `
	Role  string ` + "`" + `json:"role"` + "`" + `
}

func CoreToResponse(dataCore users.Core) Response {
	return Response{
		Id:    dataCore.Id,
		Email: dataCore.Email,
		Name:  dataCore.Name,
		Role:  dataCore.Role,
	}
}

func ListCoreToResponse(dataCore []users.Core) []Response {
	var dataResponse []Response
	for _, v := range dataCore {
		dataResponse = append(dataResponse, CoreToResponse(v))
	}
	return dataResponse
}
	
`
	return text
}
