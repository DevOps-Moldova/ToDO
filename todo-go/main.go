package main

import (
	"time"

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:4200"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/health", func(ctx *gin.Context) {
		message := "ok"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ToDoRouteController.ToDoRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
