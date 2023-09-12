package routes

import (
	docs "github.com/DevOps-Moldova/ToDo/todo-go/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocRoute(eng *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
