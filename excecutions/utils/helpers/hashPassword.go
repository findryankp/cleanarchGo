package helpers

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func hashCreate() {
	file, err := generates.FilesCreate("./utils/helpers/hashPassword.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, hashContent())
		fmt.Println("Hash File Created")
	}
}

func hashContent() string {
	var text = `package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
	`

	return text
}
