package service

import (
	"errors"

	"task-api/internal/models"
	"task-api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

var validStatuses = map[string]bool{
	"pending":     true,
	"in-progress": true,
	"completed":   true,
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) validateStatus(status string) bool {
	return validStatuses[status]
}

func (s *TaskService) CreateTask(task *models.Task) error {

	if !s.validateStatus(task.Status) {
		return errors.New("invalid status")
	}

	return s.repo.Create(task)
}

func (s *TaskService) GetTasks(
	page int,
	limit int,
	status string,
	dueDate string,
) ([]models.Task, error) {

	return s.repo.GetAll(
		page,
		limit,
		status,
		dueDate,
	)
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(id uint, updatedTask *models.Task) error {

	if !s.validateStatus(updatedTask.Status) {
		return errors.New("invalid status")
	}

	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate

	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {

	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
