package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func envCreate() {
	file, err := generates.FilesCreate("./local.env")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, envContent())
		fmt.Println("Env File Created")
	}
}

func envContent() string {
	var text = `
#mysql database
DBUSERNAME=findryankp
DBPASS=passwordfindryan
DBHOST=127.0.0.1
DBPORT=3306
DBNAME=starterpackdb

#JWT
JWTKEY=RahasiaBos	
`
	return text
}
