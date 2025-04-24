package dto

import (
	"time"
)

type ProductRequestDto struct {
	UserID      uint   `gorm:"not null"`
	Name        string `json:"name" binding:"required" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
}

type ProductDetailDto struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProductResponseDto struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint            `json:"user_id" gorm:"not null"`
	User        UserResponseDto `json:"user" gorm:"foreignKey:UserID"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Description string          `json:"description" gorm:"type:text"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}
