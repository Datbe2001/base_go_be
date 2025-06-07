package repo

import (
	"base_go_be/global"
	"base_go_be/internal/model"
	"context"

	"github.com/jackc/pgxpool/v5"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByID(id uint) *model.User
	GetListUser() []*model.User
	CreateUser(user *model.User) (int, error)
}

func NewUserRepository() IUserRepository {
	return &userRepository{db: global.Postgres}
}

type userRepository struct {
	db *pgxpool.Pool
}

func (r *userRepository) GetUserByEmail(email string) bool {
	ctx := context.Background()
	var count int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	return err == nil && count > 0
}

func (r *userRepository) GetUserByID(id uint) *model.User {
	ctx := context.Background()
	var user model.User
	err := r.db.QueryRow(ctx,
		"SELECT id, username, email, is_active, role, created_at FROM users WHERE id = $1",
		id).Scan(&user.ID, &user.Username, &user.Email, &user.IsActive, &user.Role, &user.CreatedAt)

	if err != nil {
		return nil
	}
	return &user
}

func (r *userRepository) GetListUser() []*model.User {
	ctx := context.Background()
	rows, err := r.db.Query(ctx, "SELECT id, username, email, is_active, role, created_at FROM users")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.IsActive, &user.Role, &user.CreatedAt)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}
	return users
}

func (r *userRepository) CreateUser(user *model.User) (int, error) {
	ctx := context.Background()
	var id int
	err := r.db.QueryRow(ctx,
		"INSERT INTO users (username, email, is_active, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Username, user.Email, user.IsActive, user.Role).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
