package roots

import "github.com/Findryankp/cleanarchGo/excecutions/generates"

func InitRoots(featuresName string) {
	dockerCreate(featuresName)
	gitigonreCreate(featuresName)
	envCreate(featuresName)
	loadConfigCreate(featuresName)
	mainAdd()
}

func mainAdd() {
	generates.PackageIntall("")
	line := generates.ContentGetLinenumber("./main.go", "cleanarchGo.Init()")
	generates.ContentInsertAtLinenumber("./main.go", mainContent(), line)
}

func mainContent() string {
	var text = `
	if err := cleanarchGo.Init(); err != nil {
		fmt.Println(err)
		return
	}

	//run on port 8080
	loadConfigs(":8080")
`

	return text
}
