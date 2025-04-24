package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	IsActive  bool      `gorm:"not null;default:true"`
	Role      string    `gorm:"type:enum('ADMIN', 'USER');not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (u *User) TableName() string {
	return "go_user"
}
