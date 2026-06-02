package repository

import (
	"gorm.io/gorm"

	"task-api/internal/models"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetAll(
	page int,
	limit int,
	status string,
	dueDate string,
) ([]models.Task, error) {

	var tasks []models.Task

	query := r.db.Model(&models.Task{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if dueDate != "" {
		query = query.Where("DATE(due_date) = ?", dueDate)
	}

	offset := (page - 1) * limit

	err := query.
		Limit(limit).
		Offset(offset).
		Find(&tasks).
		Error

	return tasks, err
}

func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {

	var task models.Task

	err := r.db.First(&task, id).Error

	return &task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
