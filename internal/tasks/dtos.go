package tasks

import (
	"fmt"
	"time"

	"github.com/reon150/go-todo-rest-api/internal/models"
	"github.com/reon150/go-todo-rest-api/pkg/utils"
)

// Request DTOs

type CreateTaskRequestDTO struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type UpdateTaskRequestDTO struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Response DTOs

type GetOneByIdResponseDTO struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Status    models.TaskStatus `json:"status"`
	TodoID    uint              `json:"todo_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type CreateTaskResponseDTO struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Status    models.TaskStatus `json:"status"`
	TodoID    uint              `json:"todo_id"`
	CreatedAt time.Time         `json:"created_at"`
}

type UpdateTaskResponseDTO struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Status    models.TaskStatus `json:"status"`
	TodoID    uint              `json:"todo_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// Validation Methods

func (dto *CreateTaskRequestDTO) Validate() *utils.APIErrorResponse {
	validationErr := utils.NewAPIErrorResponse()

	if dto.Title == "" {
		validationErr.AddFieldError("title", "Title is required")
	}

	if err := validateStatus(&dto.Status); err != nil {
		validationErr.AddFieldError("status", err.Error())
	}

	if validationErr.HasErrors() {
		validationErr.AddGeneralError("Validation errors occurred")
		return validationErr
	}
	return nil
}

func (dto *UpdateTaskRequestDTO) Validate() *utils.APIErrorResponse {
	validationErr := utils.NewAPIErrorResponse()

	if dto.Title == "" {
		validationErr.AddFieldError("title", "Title is required")
	}

	if err := validateStatus(&dto.Status); err != nil {
		validationErr.AddFieldError("status", err.Error())
	}

	if validationErr.HasErrors() {
		validationErr.AddGeneralError("Validation errors occurred")
		return validationErr
	}
	return nil
}

func validateStatus(status *string) error {
	var validStatuses = map[string]bool{
		string(models.TaskStatusPending):    true,
		string(models.TaskStatusInProgress): true,
		string(models.TaskStatusCompleted):  true,
	}

	if status != nil {
		if !validStatuses[*status] {
			return fmt.Errorf("the status: %s is invalid", *status)
		}
	}
	return nil
}
