package repo

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func (*userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
