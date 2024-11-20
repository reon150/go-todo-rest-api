package todos

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

type TodoHandler struct {
	service TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HTTP Method: %s\n", r.Method)
	w.Header().Set("Content-Type", "application/json")
	println("URL: %v", r.URL.Path)

	id := getIDFromPath(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if id == nil {
			h.GetTodosHandler(w, r)
		} else {
			h.GetTodoByIDHandler(w, r, id)
		}
	case http.MethodPost:
		h.CreateTodoHandler(w, r)
	case http.MethodPut:
		h.UpdateTodoHandler(w, r, id)
	case http.MethodDelete:
		h.DeleteTodoHandler(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(GetTodosResponseDTO(todos))
}

func (h *TodoHandler) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	todo, err := h.service.GetTodoByID(*id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Todo with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ToGetOneByIdResponseDTO(todo))
}

func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var dto *CreateTodoRequestDTO

	if apiError := utils.DecodeJSONBody(w, r, &dto); apiError != nil {
		return
	}

	if apiError := dto.Validate(); apiError != nil {
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	todo := ToModelFromCreateDTO(dto)

	if err := h.service.CreateTodo(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ToCreateTodoResponseDTO(todo))
}

func (h *TodoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	//TODO: Extract this validation to a new method
	if id == nil {
		apiError := utils.NewAPIErrorResponse()
		apiError.AddGeneralError("Validation errors occurred")
		apiError.AddFieldError("id", "ID is required in the URL path")
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	//TODO: Move this validation to UpdateTodo service
	if _, err := h.service.GetTodoByID(*id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Todo with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var dto *UpdateTodoRequestDTO

	if apiError := utils.DecodeJSONBody(w, r, &dto); apiError != nil {
		return
	}

	if apiError := dto.Validate(); apiError != nil {
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	todo := ToModelFromUpdateDTO(id, dto)

	if err := h.service.UpdateTodo(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ToUpdateTodoResponseDTO(todo))
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request, id *uint) {
	//TODO: Extract this validation to a new method
	if id == nil {
		apiError := utils.NewAPIErrorResponse()
		apiError.AddGeneralError("Validation errors occurred")
		apiError.AddFieldError("id", "ID is required in the URL path")
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	//TODO: Move this validation to DeleteTodo service
	if _, err := h.service.GetTodoByID(*id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiError := utils.NewAPIErrorResponse(http.StatusNotFound)
			apiError.AddGeneralError(fmt.Sprintf("Todo with ID %d was not found", *id))
			w.WriteHeader(apiError.Code)
			json.NewEncoder(w).Encode(apiError)
			return
		}

		apiError := utils.NewInternalServerError()
		w.WriteHeader(apiError.Code)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	if err := h.service.DeleteTodo(*id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getIDFromPath(path string) *uint {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return nil
	}
	id, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return nil
	}
	uintID := uint(id)
	return &uintID
}
