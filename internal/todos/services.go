package todos

import "github.com/reon150/go-todo-rest-api/internal/models"

type TodoService interface {
	GetTodos() ([]models.Todo, error)
	GetTodoByID(id uint) (*models.Todo, error)
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uint) error
}

type TodoServiceImpl struct {
	repo *TodoRepository
}

func NewTodoService(repo *TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) GetTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoServiceImpl) GetTodoByID(id uint) (*models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoServiceImpl) CreateTodo(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoServiceImpl) UpdateTodo(todo *models.Todo) error {
	return s.repo.Update(todo)
}

func (s *TodoServiceImpl) DeleteTodo(id uint) error {
	return s.repo.SoftDelete(id)
}
