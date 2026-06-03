package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"task-api/internal/dto"
	"task-api/internal/models"
	"task-api/internal/service"
)

type TaskHandler struct {
	service service.TaskServiceInterface
}

func NewTaskHandler(service service.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Creates a task with title, optional description, status, and optional due date.
// @Tags tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param task body models.Task true "Task payload"
// @Success 201 {object} dto.CreateTaskResponse
// @Failure 400 {object} dto.MessageResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateTaskResponse{
		Message: "task created successfully",
		Task: dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
		},
	})
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Returns paginated list of tasks with optional filters.
// @Tags tasks
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param status query string false "Filter by status" Enums(pending, in-progress, completed) default(pending)
// @Param due_date query string false "Filter by due date" example(2026-06-15T14:30:00Z)
// @Success 200 {array} dto.TaskResponse
// @Failure 500 {object} dto.MessageResponse
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	status := c.Query("status")
	dueDate := c.Query("due_date")

	tasks, err := h.service.GetTasks(page, limit, status, dueDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.MessageResponse{
			Message: "could not retrieve tasks",
		})
		return
	}

	var response []dto.TaskResponse
	for _, t := range tasks {
		response = append(response, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			DueDate:     t.DueDate,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetTask godoc
// @Summary Get task by ID
// @Description Returns a single task by its ID.
// @Tags tasks
// @Security BearerAuth
// @Produce json
// @Param id path int true "Task ID" example(1)
// @Success 200 {object} dto.TaskResponse
// @Failure 400 {object} dto.MessageResponse
// @Failure 404 {object} dto.MessageResponse
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: "invalid task id",
		})
		return
	}

	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.MessageResponse{
			Message: "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	})
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Updates task fields by ID.
// @Tags tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID" example(1)
// @Param task body models.Task true "Updated task data"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.MessageResponse
// @Failure 404 {object} dto.MessageResponse
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: "invalid task id",
		})
		return
	}

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.service.UpdateTask(uint(id), &task)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "task updated successfully",
	})
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Deletes a task by ID.
// @Tags tasks
// @Security BearerAuth
// @Param id path int true "Task ID" example(1)
// @Success 204 "Task deleted successfully, no content returned"
// @Failure 400 {object} dto.MessageResponse
// @Failure 404 {object} dto.MessageResponse
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: "invalid task id",
		})
		return
	}

	err = h.service.DeleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.MessageResponse{
			Message: "task not found",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
