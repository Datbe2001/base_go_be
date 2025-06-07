package model

import "time"

type User struct {
	ID        uint      `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) TableName() string {
	return "go_user"
}
