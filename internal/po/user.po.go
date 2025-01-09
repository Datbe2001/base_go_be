package po

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID       uuid.UUID `gorm:"type:varchar(36); primaryKey"`
	Username string `gorm:"column:username; not null"`
	IsActive bool   `gorm:"column:is_active; type:boolean; not null"`
	Roles    []Role `gorm:"many2many:go_user_roles;"`
}

func (r *User) TableName() string {
	return "go_user"
}
