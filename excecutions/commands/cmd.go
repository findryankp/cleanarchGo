package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Loading() {
	Clear()
	fmt.Printf("\n\n")
	fmt.Println("------------ LOADING ---------------")
	fmt.Println("------------------------------------")
}

func Feature() {
	Clear()
	fmt.Printf("\n\n")
	fmt.Println("------------- FINISH ---------------")
	fmt.Println("------------------------------------")
	fmt.Println("|   		Feature Created         |")
	fmt.Println("------------------------------------")
}

func Init() {
	Clear()
	fmt.Printf("\n\n")
	fmt.Println("------------- FINISH ---------------")
	fmt.Println("------------------------------------")
	fmt.Println("|   	Initialization Success      |")
	fmt.Println("------------------------------------")
}

func Menu() {
	Clear()
	fmt.Printf("\n\n")
	fmt.Println("-------------- MENU ----------------")
	fmt.Println("------------------------------------")
	fmt.Println(`| 1. "init" for init project       |`)
	fmt.Println(`| 2. "features" for create feature |`)
	fmt.Println(`| 3. "featwithstruct" for create 
					   feature with initial struct  |`)
	fmt.Println("------------------------------------")
}
