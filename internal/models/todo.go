package models

import (
	"time"
)

type Todo struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Title       string     `gorm:"size:255;not null"`
	Description *string    `gorm:"type:text"`
	Status      TodoStatus `gorm:"size:20;default:'Pending'"`
	Tasks       []Task     `gorm:"foreignKey:TodoID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"index"`
}
