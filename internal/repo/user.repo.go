package repo

import (
	"base_go_be/global"
	"base_go_be/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByID(id int) *model.User
	CreateUser(user *model.User) (int, error)
}

func NewUserRepository() IUserRepository {
	return &userRepository{db: global.Mysql}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUserByEmail(email string) bool {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return err == nil
}

func (r *userRepository) GetUserByID(id int) *model.User {
	var user model.User
	_ = r.db.First(&user, id)
	return &user
}
func (r *userRepository) CreateUser(user *model.User) (int, error) {
	err := r.db.Create(user).Error
	return int(user.ID), err
}
