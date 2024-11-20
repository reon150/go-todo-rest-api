package models

import (
	"time"
)

type Task struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Title     string     `gorm:"size:255;not null"`
	Status    TaskStatus `gorm:"size:20;default:'Pending'"`
	TodoID    uint       `gorm:"not null"`
	Todo      Todo       `gorm:"foreignKey:TodoID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}
