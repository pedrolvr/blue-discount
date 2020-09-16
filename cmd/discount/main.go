package main

import (
	"blue-discount/internal/app"
	"fmt"
	"log"
)

func main() {
	config, err := app.ReadConfig("app", "./config/")

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Println("config", config)
	fmt.Println("discount service")
}
