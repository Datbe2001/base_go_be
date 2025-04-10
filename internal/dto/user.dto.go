package dto

type UserRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
