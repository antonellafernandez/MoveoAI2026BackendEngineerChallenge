package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"task-api/internal/dto"
	"task-api/internal/models"
)

// fakeTaskService is a mock implementation of the TaskService interface for testing purposes
type fakeTaskService struct{}

func (f *fakeTaskService) CreateTask(task *models.Task) error {
	task.ID = 1
	return nil
}

func (f *fakeTaskService) GetTasks(page, limit int, status, dueDate string) ([]models.Task, error) {
	if status == "completed" {
		return []models.Task{
			{ID: 2, Title: "Done Task", Status: "completed"},
		}, nil
	}

	if status == "pending" {
		return []models.Task{
			{ID: 1, Title: "Pending Task", Status: "pending"},
		}, nil
	}

	return []models.Task{
		{ID: 1, Title: "Default Task", Status: "pending"},
	}, nil
}

func (f *fakeTaskService) GetTaskByID(id uint) (*models.Task, error) {
	if id == 1 {
		return &models.Task{ID: 1, Title: "Test 1", Description: "Desc", Status: "pending"}, nil
	}
	return nil, errors.New("not found")
}

func (f *fakeTaskService) UpdateTask(id uint, task *models.Task) error {
	if id != 1 {
		return errors.New("not found")
	}
	return nil
}

func (f *fakeTaskService) DeleteTask(id uint) error {
	if id != 1 {
		return errors.New("not found")
	}
	return nil
}

// setupTaskRouter initializes the Gin router with task routes and injects the fake service
func setupTaskRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	fake := &fakeTaskService{}

	h := &TaskHandler{
		service: fake,
	}

	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetTasks)
	r.GET("/tasks/:id", h.GetTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

	return r
}

// TestCreateTask tests the task creation scenario
func TestCreateTask(t *testing.T) {
	r := setupTaskRouter()

	body := models.Task{
		Title:       "Test",
		Description: "Desc",
		Status:      "pending",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

// TestGetTasks tests the retrieval of all tasks
func TestGetTasks(t *testing.T) {
	r := setupTaskRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp []dto.TaskResponse
	_ = json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp) == 0 {
		t.Errorf("expected tasks, got empty")
	}
}

// TestGetTasks_FilterStatus tests the retrieval of tasks filtered by status
func TestGetTasks_FilterStatus(t *testing.T) {
	r := setupTaskRouter()

	req, _ := http.NewRequest("GET", "/tasks?status=completed", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp []dto.TaskResponse
	_ = json.Unmarshal(w.Body.Bytes(), &resp)

	if len(resp) != 1 {
		t.Errorf("expected 1 task, got %d", len(resp))
	}

	if resp[0].Status != "completed" {
		t.Errorf("expected completed status")
	}
}

// TestGetTask tests the retrieval of a single task by ID
func TestGetTask(t *testing.T) {
	r := setupTaskRouter()

	req, _ := http.NewRequest("GET", "/tasks/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// TestUpdateTask tests the task update scenario
func TestUpdateTask(t *testing.T) {
	r := setupTaskRouter()

	body := models.Task{
		Title:  "Updated",
		Status: "pending",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// TestDeleteTask tests the task deletion scenario
func TestDeleteTask(t *testing.T) {
	r := setupTaskRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", w.Code)
	}
}
