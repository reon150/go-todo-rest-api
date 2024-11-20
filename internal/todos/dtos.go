package todos

import (
	"fmt"
	"time"

	"github.com/reon150/go-todo-rest-api/internal/models"
	"github.com/reon150/go-todo-rest-api/pkg/utils"
)

// Request DTOs

type CreateTodoRequestDTO struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}

type UpdateTodoRequestDTO struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}

// Response DTOs

type GetOneByIdResponseDTO struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Status      models.TodoStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type CreateTodoResponseDTO struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TodoStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
}

type UpdateTodoResponseDTO struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TodoStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// Validation Methods

func (dto *CreateTodoRequestDTO) Validate() *utils.APIErrorResponse {
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

func (dto *UpdateTodoRequestDTO) Validate() *utils.APIErrorResponse {
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
		string(models.TodoStatusPending):    true,
		string(models.TodoStatusInProgress): true,
		string(models.TodoStatusCompleted):  true,
	}

	if status != nil {
		if !validStatuses[*status] {
			return fmt.Errorf("the status: %s is invalid", *status)
		}
	}
	return nil
}
