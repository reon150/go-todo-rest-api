package todos

import "github.com/reon150/go-todo-rest-api/internal/models"

// To Model

func ToModelFromCreateDTO(dto *CreateTodoRequestDTO) *models.Todo {
	return &models.Todo{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.TodoStatus(dto.Status),
	}
}

func ToModelFromUpdateDTO(ID *uint, dto *UpdateTodoRequestDTO) *models.Todo {
	return &models.Todo{
		ID:          *ID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.TodoStatus(dto.Status),
	}
}

// To DTO

func ToGetOneByIdResponseDTO(todo *models.Todo) *GetOneByIdResponseDTO {
	return &GetOneByIdResponseDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

func GetTodosResponseDTO(todos []models.Todo) []*GetOneByIdResponseDTO {
	var dtos []*GetOneByIdResponseDTO
	for _, todo := range todos {
		dtos = append(dtos, ToGetOneByIdResponseDTO(&todo))
	}
	return dtos
}

func ToCreateTodoResponseDTO(todo *models.Todo) *CreateTodoResponseDTO {
	return &CreateTodoResponseDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: *todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
	}
}

func ToUpdateTodoResponseDTO(todo *models.Todo) *UpdateTodoResponseDTO {
	return &UpdateTodoResponseDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: *todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
