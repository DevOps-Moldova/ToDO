package controllers

import (
	"strings"
	"time"

	"github.com/DevOps-Moldova/ToDo/todo-go/models"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ToDoController struct {
	DB *gorm.DB
}

func NewToDoController(DB *gorm.DB) ToDoController {
	return ToDoController{DB}
}

// @BasePath /api/v1

// AddToDo godoc
// @Summary add ToDo
// @Schemes
// @Description create new ToDo
// @Tags todo
// @Accept json
// @Produce json
// @Param   todo     body    models.ToDo     true        "ToDo details"
// @Success 200 {object} models.ToDo "ok"
// @Router /todos [post]
func (pc *ToDoController) AddToDo(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(models.ToDo)
	var payload *models.ToDo

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.ToDo{
		Name:        payload.Name,
		Description: payload.Description,
		Status:      "New",
		// Status:      payload.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"status": "fail", "message": "ToDo with that ID already exists"})
			return
		}
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

// GetToDos godoc
// @Summary get ToDOs
// @Schemes
// @Description get all ToDos
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} models.ToDo "ok"
// @Router /todos [get]
func (pc *ToDoController) GetToDos(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var todos []models.ToDo
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&todos)
	if results.Error != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "success", "results": len(todos), "data": todos})
}

// UpdateToDo godoc
// @Summary Update ToDo
// @Schemes
// @Description Update existing ToDo
// @Tags todo
// @Accept json
// @Produce json
// @Param   id     path    string     true        "ToDo ID"
// @Param   todo     body    models.ToDo     true        "ToDo details"
// @Success 200 {object} models.ToDo "ok"
// @Router /todos/{id} [put]
func (pc *ToDoController) UpdateToDo(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload *models.ToDo
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedToDo models.ToDo
	result := pc.DB.First(&updatedToDo, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	toDoToUpdate := models.ToDo{
		Name:        payload.Name,
		Description: payload.Description,
		Status:      payload.Status,
		CreatedAt:   updatedToDo.CreatedAt,
		UpdatedAt:   now,
	}

	pc.DB.Model(&updatedToDo).Updates(toDoToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedToDo})
}

// FindToDoById godoc
// @Summary Get ToDo
// @Schemes
// @Description Get existing ToDo
// @Tags todo
// @Accept json
// @Produce json
// @Param   id     path    string     true        "ToDo ID"
// @Success 200 {object} models.ToDo "ok"
// @Router /todos/{id} [get]
func (pc *ToDoController) FindToDoById(ctx *gin.Context) {
	todoId := ctx.Param("id")

	var todo models.ToDo
	result := pc.DB.First(&todo, "id = ?", todoId)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No ToDo with that ID exists"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "success", "data": todo})
}

// DeleteToDo godoc
// @Summary Delete ToDo
// @Schemes
// @Description Delete existing ToDo
// @Tags todo
// @Accept json
// @Produce json
// @Param   id     path    string     true        "ToDo ID"
// @Success 204 {object} string "ok"
// @Router /todos/{id} [delete]
func (pc *ToDoController) DeleteToDo(ctx *gin.Context) {
	todoId := ctx.Param("id")

	result := pc.DB.Delete(&models.ToDo{}, "id = ?", todoId)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No ToDo with that ID exists"})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)
}
