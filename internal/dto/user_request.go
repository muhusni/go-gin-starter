package dto

type CreateUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type UpdateUserRequest struct {
	Name            *string `json:"name"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"password_confirm"`
}
