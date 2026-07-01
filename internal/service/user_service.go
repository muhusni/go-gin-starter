package service

import (
	"errors"

	"github.com/muhusni/go-gin-starter/internal/dto"
	"github.com/muhusni/go-gin-starter/internal/model"
	"github.com/muhusni/go-gin-starter/internal/security"
	"gorm.io/gorm"
)

var ErrPasswordMismatch = errors.New("password confirm does not match")
var ErrNameRequired = errors.New("name cannot be blank")
var ErrPasswordRequired = errors.New("password cannot be blank")

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUsers() ([]dto.UserResponse, error) {
	var users []model.User

	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}

	return responses, nil
}

func (s *UserService) GetUser(id uint64) (dto.UserResponse, error) {
	var user model.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return dto.UserResponse{}, err
	}

	return ToUserResponse(user), nil
}

func (s *UserService) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	if req.Password != req.PasswordConfirm {
		return dto.UserResponse{}, ErrPasswordMismatch
	}

	passwordHash, err := security.HashPassword(req.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		IsAdmin:  false,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return dto.UserResponse{}, err
	}

	return ToUserResponse(user), nil
}

func (s *UserService) UpdateUser(id uint64, req dto.UpdateUserRequest) (dto.UserResponse, error) {

	var user model.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return dto.UserResponse{}, err
	}

	if err := applyUserUpdate(&user, &req); err != nil {
		return dto.UserResponse{}, err
	}

	if err := s.DB.Save(&user).Error; err != nil {
		return dto.UserResponse{}, err
	}

	return ToUserResponse(user), nil
}
func (s *UserService) DeleteUser(id uint64) error {
	var user model.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := s.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
func ToUserResponse(user model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func applyUserUpdate(user *model.User, req *dto.UpdateUserRequest) error {
	if req.Name != nil {
		if *req.Name == "" {
			return ErrNameRequired
		}
		user.Name = *req.Name
	}

	if req.Password == nil && req.PasswordConfirm != nil {
		return ErrPasswordRequired
	}

	if req.Password != nil {
		if *req.Password == "" {
			return ErrPasswordRequired
		}
		if req.PasswordConfirm == nil || *req.Password != *req.PasswordConfirm {
			return ErrPasswordMismatch
		}
		passwordHash, err := security.HashPassword(*req.Password)
		if err != nil {
			return err
		}
		user.Password = passwordHash
	}
	return nil
}
