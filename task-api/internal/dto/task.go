package dto

import "time"

type TaskResponse struct {
	ID          uint       `json:"id" example:"1"`
	Title       string     `json:"title" example:"Configure production server"`
	Description string     `json:"description" example:"Install dependencies and set up Nginx reverse proxy on the server"`
	Status      string     `json:"status" example:"pending"`
	DueDate     *time.Time `json:"due_date" example:"2026-06-15T14:30:00Z"`
}

type CreateTaskResponse struct {
	Message string       `json:"message" example:"task created successfully"`
	Task    TaskResponse `json:"task"`
}

type MessageResponse struct {
	Message string `json:"message" example:"response message detail"`
}
