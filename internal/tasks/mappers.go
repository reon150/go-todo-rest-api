package tasks

import "github.com/reon150/go-todo-rest-api/internal/models"

// To Model

func ToModelFromCreateDTO(TodoID *uint, dto *CreateTaskRequestDTO) *models.Task {
	return &models.Task{
		Title:  dto.Title,
		Status: models.TaskStatus(dto.Status),
		TodoID: *TodoID,
	}
}

func ToModelFromUpdateDTO(ID *uint, dto *UpdateTaskRequestDTO) *models.Task {
	return &models.Task{
		ID:     *ID,
		Title:  dto.Title,
		Status: models.TaskStatus(dto.Status),
	}
}

// To DTO

func ToGetOneByIdResponseDTO(task *models.Task) *GetOneByIdResponseDTO {
	return &GetOneByIdResponseDTO{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		TodoID:    task.TodoID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func GetTasksResponseDTO(tasks []models.Task) []*GetOneByIdResponseDTO {
	var dtos []*GetOneByIdResponseDTO
	for _, task := range tasks {
		dtos = append(dtos, ToGetOneByIdResponseDTO(&task))
	}
	return dtos
}

func ToCreateTaskResponseDTO(task *models.Task) *CreateTaskResponseDTO {
	return &CreateTaskResponseDTO{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		TodoID:    task.TodoID,
		CreatedAt: task.CreatedAt,
	}
}

func ToUpdateTaskResponseDTO(task *models.Task) *UpdateTaskResponseDTO {
	return &UpdateTaskResponseDTO{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		TodoID:    task.TodoID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}
