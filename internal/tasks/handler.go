package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/reon150/go-todo-rest-api/pkg/utils"
	"gorm.io/gorm"
)

type TaskHandler struct {
	service TaskService
}

func NewTaskHandler(service TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HTTP Method: %s\n", r.Method)
	w.Header().Set("Content-Type", "application/json")
	println("URL: %v", r.URL.Path)

	id := h.getTaskIDFromPath(r.URL.Path)
	todoID := h.getTodoIDFromPath(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if id == nil {
			h.GetTasksHandler(w, r, todoID)
		} else {
			h.GetTaskByIDHandler(w, r, id)
		}
	case http.MethodPost:
		h.CreateTaskHandler(w, r, todoID)
	case http.MethodPut:
		h.UpdateTaskHandler(w, r, id)
	case http.MethodDelete:
		h.DeleteTaskHandler(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) GetTasksHandler(w http.ResponseWriter, r *http.Request, todoID *uint) {
	tasks, err := h.service.GetTasksByTodoID(*todoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //TODO: Use APIError
		return
	}
	json.NewEncoder(w).Encode(GetTasksResponseDTO(tasks))
}

func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	task, err := h.service.GetTaskByID(*id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Task with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ToGetOneByIdResponseDTO(task))
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request, TodoID *uint) {
	//TODO: Add validation for Todo exists
	var dto *CreateTaskRequestDTO

	if apiError := utils.DecodeJSONBody(w, r, &dto); apiError != nil {
		return
	}

	if apiError := dto.Validate(); apiError != nil {
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	task := ToModelFromCreateDTO(TodoID, dto)

	if err := h.service.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //TODO: Use API error
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ToCreateTaskResponseDTO(task))
}

func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	if id == nil {
		apiError := utils.NewAPIErrorResponse()
		apiError.AddGeneralError("Validation errors occurred")
		apiError.AddFieldError("id", "ID is required in the URL path")
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	if _, err := h.service.GetTaskByID(*id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Task with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var dto *UpdateTaskRequestDTO

	if apiError := utils.DecodeJSONBody(w, r, &dto); apiError != nil {
		return
	}

	if apiError := dto.Validate(); apiError != nil {
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	task := ToModelFromUpdateDTO(id, dto)

	if err := h.service.UpdateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ToUpdateTaskResponseDTO(task))
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	if id == nil {
		apiError := utils.NewAPIErrorResponse()
		apiError.AddGeneralError("Validation errors occurred")
		apiError.AddFieldError("id", "ID is required in the URL path")
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	if _, err := h.service.GetTaskByID(*id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Task with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}

		apiError := utils.NewInternalServerError()
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	if err := h.service.DeleteTask(*id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) getTodoIDFromPath(path string) *uint {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) >= 2 && segments[0] == "todo" {
		id, err := strconv.ParseUint(segments[1], 10, 32)
		if err != nil {
			return nil
		}
		uintID := uint(id)
		return &uintID
	}
	return nil
}

func (h *TaskHandler) getTaskIDFromPath(path string) *uint {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) >= 4 && segments[0] == "todo" && segments[2] == "task" {
		id, err := strconv.ParseUint(segments[3], 10, 32)
		if err != nil {
			return nil
		}
		uintID := uint(id)
		return &uintID
	}
	return nil
}
