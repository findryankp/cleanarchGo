package main

import "fmt"

func main() {
	if err := Init(); err != nil {
		fmt.Println(err)
		return
	}
}
