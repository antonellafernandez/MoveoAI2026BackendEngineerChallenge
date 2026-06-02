package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"task-api/internal/database"
	"task-api/internal/handlers"
	"task-api/internal/middleware"
	"task-api/internal/repository"
	"task-api/internal/service"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(repo)

	taskHandler := handlers.NewTaskHandler(taskService)

	router := gin.Default()

	// Ruta pública
	router.POST("/login", handlers.Login)

	// Rutas protegidas
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/tasks", taskHandler.CreateTask)
	protected.GET("/tasks", taskHandler.GetTasks)
	protected.GET("/tasks/:id", taskHandler.GetTask)
	protected.PUT("/tasks/:id", taskHandler.UpdateTask)
	protected.DELETE("/tasks/:id", taskHandler.DeleteTask)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
