package main

import (
	"log"

	"github.com/DevOps-Moldova/ToDo/todo-go/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.RunMigrations()
}
