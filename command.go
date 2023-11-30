package cleanarchGo

import (
	"errors"
	"fmt"
	"os"

	"github.com/Findryankp/cleanarchGo/excecutions"
	"github.com/Findryankp/cleanarchGo/excecutions/commands"
)

func Init() error {
	var argsRaw = os.Args
	if len(argsRaw) < 2 {
		return nil
	}

	args := map[string]bool{
		"init":           CommandInit(),
		"features":       CommandFeature(argsRaw),
		"featwithstruct": CommandInit(),
	}

	if v, ok := args[argsRaw[1]]; ok {
		if v {
			return errors.New("run command finish")
		}
	} else {
		commands.Menu()
	}

	return errors.New("run command finish")

}

func CommandInit() bool {
	commands.Loading()
	excecutions.CommandInit()
	commands.Init()

	return true
}

func CommandFeature(arg []string) bool {
	if len(arg) != 3 {
		fmt.Println("[COMMAND] - wrong argument (ex: go run . features books)")
		return false
	}

	if arg[2][0] >= 97 && arg[2][0] <= 122 {
		commands.Loading()
		excecutions.NewFeatures(arg[2])
		commands.Feature()
		return true
	}

	fmt.Println("[COMMAND] - make feature with lowercase first (ex: go run . features books)")

	return false
}

func CommandFeatureWithStruct(features string, structString string) bool {
	commands.Loading()
	excecutions.NewFeaturesWithStruct(features, structString)
	commands.Feature()

	return false
}

var data = `
	type Books struct {
		Id    uint
		Class string
	}
`