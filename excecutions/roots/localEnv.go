package roots

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func envCreate(featuresName string) {
	base := "./" + featuresName
	file, err := generates.FilesCreate(base + "/local.env")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, envContent(featuresName))
		fmt.Println("Env File Created")
	}
}

func envContent(featuresName string) string {
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
