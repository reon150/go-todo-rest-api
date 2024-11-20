package migrations

import (
	"github.com/reon150/go-todo-rest-api/internal/models"
	"gorm.io/gorm"
)

func CreateTodoTable(db *gorm.DB) error {
	m := db.Migrator()

	if !m.HasTable(&models.Todo{}) {
		if err := db.Exec(`CREATE TABLE todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT
		)`).Error; err != nil {
			return err
		}

		if err := m.AddColumn(&models.Todo{}, "title"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Todo{}, "description"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Todo{}, "status"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Todo{}, "created_at"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Todo{}, "updated_at"); err != nil {
			return err
		}
		if err := m.AddColumn(&models.Todo{}, "deleted_at"); err != nil {
			return err
		}

		if err := m.CreateIndex(&models.Todo{}, "idx_todos_deleted_at"); err != nil {
			return err
		}
	}

	return nil
}

func RollbackTodoTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Todo{})
}
