package models

import (
	"fmt"
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

func (task *Task) BeforeCreate(tx *gorm.DB) error {
	if err := ValidateTaskStatus(task.Status); err != nil {
		return err
	}
	return nil
}

func (task *Task) BeforeUpdate(tx *gorm.DB) error {
	if err := ValidateTaskStatus(task.Status); err != nil {
		return err
	}
	return nil
}

func ValidateTaskStatus(status bool) error {
	if status != true && status != false {
		return fmt.Errorf("Status must be true or false")
	}
	return nil
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
