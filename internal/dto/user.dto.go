package dto

type UserRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserResponseDto struct {
	Id       uint   `json:"id" binding:"required"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
