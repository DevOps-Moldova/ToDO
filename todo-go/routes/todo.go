package routes

import (
	"github.com/DevOps-Moldova/ToDo/todo-go/controllers"

	"github.com/gin-gonic/gin"
)

type ToDoRouteController struct {
	toDoController controllers.ToDoController
}

func NewRouteToDoController(todoController controllers.ToDoController) ToDoRouteController {
	return ToDoRouteController{todoController}
}

func (pc *ToDoRouteController) ToDoRoute(rg *gin.RouterGroup) {

	router := rg.Group("todos")

	router.POST("/", pc.toDoController.AddToDo)
	router.GET("/", pc.toDoController.GetToDos)
	router.PUT("/:id", pc.toDoController.UpdateToDo)
	router.GET("/:id", pc.toDoController.FindToDoById)
	router.DELETE("/:id", pc.toDoController.DeleteToDo)
}
