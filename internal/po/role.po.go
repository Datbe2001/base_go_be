package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	//ID       uuid.UUID `gorm:"column:id; type:int; primary_key; not null; default:uuid_generate_v4()"`
	RoleName string `gorm:"column:role_name; not null"`
}

func (r *Role) TableName() string {
	return "go_role"
}
