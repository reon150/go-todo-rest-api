package tasks

import (
	"time"

	"github.com/reon150/go-todo-rest-api/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("deleted_at IS NULL").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByTodoID(todoID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("todo_id = ? AND deleted_at IS NULL", todoID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) SoftDelete(id uint) error {
	now := time.Now()
	err := r.db.Model(&models.Task{}).Where(&models.Task{ID: id}).Update("deleted_at", &now).Error
	return err
}

func (r *TaskRepository) SoftDeleteByTodoID(todoID uint) error {
	now := time.Now()
	err := r.db.Model(&models.Task{}).Where(&models.Task{TodoID: todoID}).Update("deleted_at", &now).Error
	return err
}
