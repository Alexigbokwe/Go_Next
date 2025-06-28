package dto

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	FullName string `json:"full_name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
