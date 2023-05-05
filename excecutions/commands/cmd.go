package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Loading() {
	fmt.Printf("\n\n")
	fmt.Println("------------------------------------")
	fmt.Println("|   Loading .....                  |")
	fmt.Println("------------------------------------")
}

func Feature() {
	fmt.Printf("\n\n")
	fmt.Println("------------------------------------")
	fmt.Println("|   Feature Created                |")
	fmt.Println("------------------------------------")
}

func Init() {
	fmt.Printf("\n\n")
	fmt.Println("------------------------------------")
	fmt.Println("|   Initialization Success         |")
	fmt.Println("------------------------------------")
}
