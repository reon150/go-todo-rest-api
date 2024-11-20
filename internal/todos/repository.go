package todos

import (
	"time"

	"github.com/reon150/go-todo-rest-api/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.db.Where("deleted_at IS NULL").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) SoftDelete(id uint) error {
	now := time.Now()
	err := r.db.Model(&models.Todo{}).Where(&models.Todo{ID: id}).Update("deleted_at", &now).Error
	return err
}
