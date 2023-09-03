package main

import (
	"fmt"

	"github.com/arferreira/goham/configs"
)

func main() {
	config, err := configs.GetConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.DBDriver)

	fmt.Println("Hello, I am Goham!")
}
