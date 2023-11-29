package models

import (
	"errors"
	"fmt"

	"github.com/Bobby-P-dev/FinalProject3_kel7/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"omitempty,min=6" `
	Role     string `json:"role" gorm:"not null" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	u.Role = "member"

	if err := validator.New().Struct(u); err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {

	if u.Role != "admin" && u.Role != "member" {
		return errors.New("invalid role")
	}

	if err := validator.New().StructExcept(u, "Password"); err != nil {
		return err
	}

	if u.Password != "" {
		if err := u.validatePassword(); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("User updated successfully")

	return nil
}

func (u *User) validatePassword() error {
	if len(u.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters long")
	}
	return nil
}

type UsersRespon struct {
	ID       int    `json:"id" form:"id"`
	FullName string `json:"Full_name" form:"full_name" `
	Email    string `json:"email" form:"email"`
}

func (UsersRespon) TableName() string {
	return "users"
}
