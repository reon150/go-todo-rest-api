package migrations

import (
	"github.com/reon150/go-todo-rest-api/internal/models"
	"gorm.io/gorm"
)

func CreateTaskTable(db *gorm.DB) error {
	m := db.Migrator()

	if !m.HasTable(&models.Task{}) {
		if err := db.Exec(`CREATE TABLE tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT
		)`).Error; err != nil {
			return err
		}

		if err := m.AddColumn(&models.Task{}, "title"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Task{}, "status"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Task{}, "todo_id"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Task{}, "created_at"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Task{}, "updated_at"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Task{}, "deleted_at"); err != nil {
			return err
		}

		if err := m.CreateIndex(&models.Task{}, "idx_tasks_deleted_at"); err != nil {
			return err
		}
	}

	return nil
}

func RollbackTaskTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Task{})
}
