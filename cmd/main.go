package main

import (
	"github.com/NorousS/CRUD/internal/handlers"
	"github.com/NorousS/CRUD/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()
	defer models.CloseDB()

	r := gin.Default()

	r.GET("/", handlers.GetHandler)
	r.GET("/:id", handlers.GetHandlerByID)
	r.POST("/", handlers.PostHandler)
	r.PATCH("/:id", handlers.UpdateHandler)
	r.DELETE("/:id", handlers.DeleteHandler)

	r.Run(":8080")
}
