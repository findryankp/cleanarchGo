package generates

import (
	"fmt"
	"os"
)

func FolderCreateList(folderName []string) {
	for _, v := range folderName {
		FolderCreate(v)
	}
}

func FolderCreate(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err = os.Mkdir(folderName, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func FolderInit() {
	folderName := []string{
		"apps",
		"features",
		"utils",
		"apps/configs",
		"apps/database",
		"apps/middlewares",
		"apps/routes",
		"utils/helpers",
	}
	FolderCreateList(folderName)
}
