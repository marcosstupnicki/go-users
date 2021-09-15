package users

import (
	"time"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        int      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        int `gorm:"primarykey"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
