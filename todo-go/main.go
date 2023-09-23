package main

import (
	"github.com/DevOps-Moldova/ToDo/todo-go/controllers"
	"github.com/DevOps-Moldova/ToDo/todo-go/initializers"
	"github.com/DevOps-Moldova/ToDo/todo-go/routes"

	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	ToDoController      controllers.ToDoController
	ToDoRouteController routes.ToDoRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	ToDoController = controllers.NewToDoController(initializers.DB)
	ToDoRouteController = routes.NewRouteToDoController(ToDoController)

	server = gin.Default()
	config_cors := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true

	server.Use(cors.New(config_cors))
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	if config.MigrateOnStart {
		initializers.RunMigrations()
	}

	routes.DocRoute(server)
	router := server.Group("/api")

	router.GET("/health", func(ctx *gin.Context) {
		message := "ok"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ToDoRouteController.ToDoRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
