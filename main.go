package main

import "fmt"

func main() {
	// Init()
	if err := Init(); err != nil {
		fmt.Println(err)
		return
	}

	//run on port 8080
	// loadConfigs(":8080")
}
