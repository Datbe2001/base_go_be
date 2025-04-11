package dto

import (
	"time"
)

type ProductRequestDto struct {
	UserID      uint   `gorm:"not null"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
}

type ProductDetailDto struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type ProductResponseDto struct {
	ID          uint            `gorm:"primaryKey;autoIncrement"`
	UserID      uint            `gorm:"not null"`
	User        UserResponseDto `gorm:"foreignKey:UserID"`
	Name        string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}
