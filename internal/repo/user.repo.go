package repo

import (
	"base_go_be/global"
	"base_go_be/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByEmailFull(email string) *model.User
	GetUserByID(id uint) *model.User
	GetListUser() []*model.User
	CreateUser(user *model.User) (uint, error)
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

func (r *userRepository) GetUserByEmailFull(email string) *model.User {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

func (r *userRepository) GetUserByID(id uint) *model.User {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil
	}
	return &user
}

func (r *userRepository) GetListUser() []*model.User {
	var users []*model.User
	r.db.Find(&users)
	return users
}

func (r *userRepository) CreateUser(user *model.User) (uint, error) {
	err := r.db.Create(user).Error
	return user.ID, err
}
