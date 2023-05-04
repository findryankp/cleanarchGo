package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Findryankp/cleanarchGo/excecutions"
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
		fmt.Println(`1. "init" for init project`)
		fmt.Println(`2. "features or -f" for create feature`)
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
		excecutions.NewFeatures(arg[2])
	}
}

func CommandInit() {
	excecutions.CommandInit()
}
