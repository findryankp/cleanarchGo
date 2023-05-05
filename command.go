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
	if len(argsRaw) >= 2 {
		Command(argsRaw)
		return errors.New("run command finish")
	}

	return nil
}

func Command(argsRaw []string) {
	args := []string{
		"features", "init",
	}

	flag := false
	for _, v := range args {
		if v == argsRaw[1] {
			ExcecutionArgument(argsRaw)
			flag = true
		}
	}

	if !flag {
		commands.Menu()
	}
}

func ExcecutionArgument(arg []string) {
	if arg[1] == "features" {
		CommandFeature(arg)
	} else if arg[1] == "init" {
		CommandInit()
	}
}

func CommandFeature(arg []string) {
	if len(arg) != 3 {
		fmt.Println("wrong argument")
	} else {
		if arg[2][0] >= 97 && arg[2][0] <= 122 {
			commands.Loading()
			excecutions.NewFeatures(arg[2])
			commands.Feature()
		} else {
			fmt.Println("make feature with lowercase first")
		}
	}
}

func CommandInit() {
	commands.Loading()
	excecutions.CommandInit()
	commands.Init()
}
