package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func dockerCreate() {
	file, err := generates.FilesCreate("./Dockerfile")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, dockerContent())
		fmt.Println("Docker File Created")
	}
}

func dockerContent() string {
	var text = `FROM golang:1.20-alpine

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

EXPOSE 8080

# create executable
RUN go build -o ` + generates.ModuleName + `

# menjalankan file executablenya
CMD ["./` + generates.ModuleName + `"]
	`

	return text
}
