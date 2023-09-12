package initializers

import (
	"github.com/DevOps-Moldova/ToDo/todo-go/models"

	"fmt"
	"log"
)

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ConnectDB(&config)
}

func RunMigrations() {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(&models.ToDo{})
	fmt.Println("Migration complete")
}

// func main() {
// 	RunMigrations()
// }
