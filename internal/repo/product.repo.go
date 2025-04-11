package repo

import (
	"base_go_be/global"
	"base_go_be/internal/model"
	"errors"
	"gorm.io/gorm"
)

//type IUserRepository interface {
//	GetUserByEmail(email string) bool
//	GetUserByID(id uint) *model.User
//	GetListUser() []*model.User
//	CreateUser(user *model.User) (int, error)
//}
//
//func NewUserRepository() IUserRepository {
//	return &userRepository{db: global.Mysql}
//}
//
//type userRepository struct {
//	db *gorm.DB
//}

type IProductRepository interface {
	FindByID(id uint) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Create(product *model.Product) (*model.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{db: global.Mysql}
}

func (pr *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := pr.db.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	if err := pr.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
