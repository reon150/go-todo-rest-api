package tasks

import "github.com/reon150/go-todo-rest-api/internal/models"

type TaskService interface {
	GetTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	GetTasksByTodoID(todoID uint) ([]models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
	DeleteTasksByTodoID(todoID uint) error
}

type TaskServiceImpl struct {
	repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) GetTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskServiceImpl) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskServiceImpl) GetTasksByTodoID(todoID uint) ([]models.Task, error) {
	return s.repo.GetByTodoID(todoID)
}

func (s *TaskServiceImpl) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *TaskServiceImpl) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskServiceImpl) DeleteTask(id uint) error {
	return s.repo.SoftDelete(id)
}

func (s *TaskServiceImpl) DeleteTasksByTodoID(todoID uint) error {
	return s.repo.SoftDeleteByTodoID(todoID)
}
