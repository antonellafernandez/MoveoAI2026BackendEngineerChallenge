package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "task-api/docs"

	"task-api/internal/config"
	"task-api/internal/database"
	"task-api/internal/handlers"
	"task-api/internal/middleware"
	"task-api/internal/repository"
	"task-api/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task API
// @version 1.0
// @description RESTful API for a simple task management system secured with JWT authentication. The API allows users to create, read, update, and delete tasks. Each task includes a title, optional description, status, and optional due date.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.Frontend.URL},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public route
	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, cfg)
	})

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))

	protected.POST("/tasks", taskHandler.CreateTask)
	protected.GET("/tasks", taskHandler.GetTasks)
	protected.GET("/tasks/:id", taskHandler.GetTask)
	protected.PUT("/tasks/:id", taskHandler.UpdateTask)
	protected.DELETE("/tasks/:id", taskHandler.DeleteTask)

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
