package service

import "task-api/internal/models"

type TaskServiceInterface interface {
	CreateTask(task *models.Task) error
	GetTasks(page, limit int, status, dueDate string) ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(id uint, task *models.Task) error
	DeleteTask(id uint) error
}
