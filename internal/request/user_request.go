package request

type UserRequest struct {
	Name            string		`json:"name" binding:"required"`
	Username        string		`json:"username" binding:"required"`
	Email           string		`json:"email" binding:"required"`
	Password        string		`json:"password" binding:"required"`
	PasswordConfirm string		`json:"password_confim" binding:"required"`
	IsAdmin         bool  		`json:"is_admin" binding:"required"`
}