package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Status      bool   `gorm:"not null" validate:"required,boolean" json:"status"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	CategoryID  uint   `gorm:"not null" json:"category_id"`
	User        UsersRespon
}

type TaskRespon struct {
	ID          int       `json:"id" form:"id"`
	Title       string    ` json:"title"`
	Description string    ` json:"description"`
	UserID      uint      ` json:"user_id"`
	CategoryID  uint      ` json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

func (TaskRespon) TableName() string {
	return "tasks"
}
