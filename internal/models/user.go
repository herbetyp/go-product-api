package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Username    string         `json:"username,omitempty" gorm:"not null"`
	Email       string         `json:"email,omitempty" gorm:"unique;not null"`
	Password    string         `json:"password,omitempty" gorm:"not null"`
	IsAdmin     bool           `json:"is_admin,omitempty" gorm:"default:false"`
	Active      bool           `json:"active" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	LastLoginAt time.Time      `json:"last_login_at,omitempty"`
}

type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(username string, email string, passw string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: passw,
	}
}

func NewUserWithID(id uint, username string, passw string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: passw,
	}
}

func FilterUserResult(u User) *User {
	return &User{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		LastLoginAt: u.LastLoginAt,
		IsAdmin:     u.IsAdmin,
		Active:      u.Active,
	}
}
