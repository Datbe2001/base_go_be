package repo

import (
	"base_go_be/global"
	"base_go_be/internal/model"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxpool/v5"
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
	db *pgxpool.Pool
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{db: global.Postgres}
}

func (pr *ProductRepository) FindByID(id uint) (*model.Product, error) {
	ctx := context.Background()
	var product model.Product
	err := pr.db.QueryRow(ctx,
		`SELECT p.id, p.user_id, p.name, p.description, p.created_at, p.updated_at,
		 u.id, u.username, u.email, u.is_active, u.role, u.created_at
		 FROM products p 
		 LEFT JOIN users u ON p.user_id = u.id 
		 WHERE p.id = $1`, id).Scan(
		&product.ID, &product.UserID, &product.Name, &product.Description,
		&product.CreatedAt, &product.UpdatedAt,
		&product.User.ID, &product.User.Username, &product.User.Email,
		&product.User.IsActive, &product.User.Role, &product.User.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) FindAll() ([]model.Product, error) {
	ctx := context.Background()
	rows, err := pr.db.Query(ctx,
		`SELECT p.id, p.user_id, p.name, p.description, p.created_at, p.updated_at,
		 u.id, u.username, u.email, u.is_active, u.role, u.created_at
		 FROM products p 
		 LEFT JOIN users u ON p.user_id = u.id`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.ID, &product.UserID, &product.Name, &product.Description,
			&product.CreatedAt, &product.UpdatedAt,
			&product.User.ID, &product.User.Username, &product.User.Email,
			&product.User.IsActive, &product.User.Role, &product.User.CreatedAt)
		if err != nil {
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	ctx := context.Background()
	var id int
	err := pr.db.QueryRow(ctx,
		"INSERT INTO products (user_id, name, description) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		product.UserID, product.Name, product.Description).Scan(&id, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}
	product.ID = uint(id)
	return product, nil
}
