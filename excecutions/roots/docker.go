package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func dockerCreate(featuresName string) {
	base := "./" + featuresName
	file, err := generates.FilesCreate(base + "/Dockerfile")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, dockerContent(featuresName))
		fmt.Println("Config File Created")
	}
}

func dockerContent(featuresName string) string {
	var text = `FROM golang:1.20-alpine

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

EXPOSE 8080

# create executable
RUN go build -o ` + featuresName + `

# menjalankan file executablenya
CMD ["./` + featuresName + `"]
	`

	return text
}
