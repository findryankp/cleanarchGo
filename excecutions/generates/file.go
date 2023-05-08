package generates

import (
	"errors"
	"fmt"
	"os"
)

func FilesCreate(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return file, errors.New("file already exists")
		} else {
			return file, errors.New(err.Error())
		}
	}
	return file, nil
}

func FileCreateList(fileNames []string) {
	for _, v := range fileNames {
		FilesCreate(v)
	}
}

func FilesAddContent(file *os.File, content string) {
	file.WriteString(content)
}

func FilesDelete(filename string) {
	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)
	}

	fmt.Println("file does'n exist")
}
