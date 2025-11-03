package handlers

import (
	"net/http"
	"strconv"

	"github.com/NorousS/CRUD/internal/models"
	"github.com/NorousS/CRUD/internal/storage"
	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	todos, err := storage.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func GetHandlerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
		return
	}
	todo, err := storage.GetTodoByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not find todos with this ID"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func PostHandler(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Bad request"})
		return
	}
	if todo.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if err := storage.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Database error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task was added"})
}

func UpdateHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
		return
	}
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Bad request"})
		return
	}
	if todo.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if err := storage.UpdateTodo(id, &todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not find todos with this ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task was updated"})
}

func DeleteHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad ID"})
		return
	}
	if err := storage.DeleteTodo(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not find todos with this ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task was deleted"})
}
