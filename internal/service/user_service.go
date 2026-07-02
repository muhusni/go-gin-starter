package service

import (
	"errors"

	"github.com/muhusni/go-gin-starter/internal/dto"
	"github.com/muhusni/go-gin-starter/internal/model"
	"github.com/muhusni/go-gin-starter/internal/repository"
	"github.com/muhusni/go-gin-starter/internal/security"
)

var ErrPasswordMismatch = errors.New("password confirm does not match")
var ErrNameRequired = errors.New("name cannot be blank")
var ErrPasswordRequired = errors.New("password cannot be blank")

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepository}
}

func (s *UserService) GetUsers() ([]dto.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}

	return responses, nil
}

func (s *UserService) GetUser(id uint64) (dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
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

	if err := s.userRepo.Create(&user); err != nil {
		return dto.UserResponse{}, err
	}

	return ToUserResponse(&user), nil
}

func (s *UserService) UpdateUser(id uint64, req dto.UpdateUserRequest) (dto.UserResponse, error) {

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if err := applyUserUpdate(user, &req); err != nil {
		return dto.UserResponse{}, err
	}

	if err := s.userRepo.Update(user); err != nil {
		return dto.UserResponse{}, err
	}

	return ToUserResponse(user), nil
}
func (s *UserService) DeleteUser(id uint64) error {
	return s.userRepo.Delete(id)
}
func ToUserResponse(user *model.User) dto.UserResponse {
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
