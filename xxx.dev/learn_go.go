package main

import "fmt"

type IUserRepo interface{ GetByID(id uint) bool }
type IOrderService interface {
	Create(userID uint, amt float64) error
}

type userRepo struct {
}

func NewUserRepo() IUserRepo { return &userRepo{} }

func (r *userRepo) GetByID(id uint) bool { return id == 1 }

type orderService struct {
	repo IUserRepo
}

func NewOrderService(r IUserRepo) IOrderService {
	return &orderService{repo: r}
}
func (s *orderService) Create(userID uint, amt float64) error {
	if !s.repo.GetByID(userID) {
		return fmt.Errorf("user %d không tồn tại", userID)
	}
	fmt.Println("Tạo order cho user", userID, "số tiền", amt)
	return nil
}

func main() {
	repo := NewUserRepo()
	svc := NewOrderService(repo) // tiêm repo vào service
	_ = svc.Create(1, 50.0)      // lỗi: user 2 không tồn tại
}
