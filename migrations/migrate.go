package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&Migration{}); err != nil {
		return err
	}

	migrations := []struct {
		name string
		fn   func(*gorm.DB) error
	}{
		{"create_todo_table", CreateTodoTable},
		{"create_task_table", CreateTaskTable},
	}

	for _, m := range migrations {
		if db.Where("name = ?", m.name).First(&Migration{}).RowsAffected == 0 {
			if err := m.fn(db); err != nil {
				return err
			}

			if err := db.Create(&Migration{Name: m.name}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
