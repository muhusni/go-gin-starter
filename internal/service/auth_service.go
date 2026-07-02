package service

import (
	"errors"

	"github.com/muhusni/go-gin-starter/internal/dto"
	"github.com/muhusni/go-gin-starter/internal/model"
	"github.com/muhusni/go-gin-starter/internal/repository"
	"github.com/muhusni/go-gin-starter/internal/security"
)

type AuthService struct {
	jwt      *security.JWTService
	userRepo *repository.UserRepository
}

var ErrInvalidCredentials = errors.New("invalid email or password")

func NewAuthService(jwt *security.JWTService, userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		jwt:      jwt,
		userRepo: userRepository,
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !(security.CheckPasswordHash(user.Password, req.Password)) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwt.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password confirm does not match")
	}

	passwordHash, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		IsAdmin:  false,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateToken(user.ID, user.Email, user.IsAdmin)
	return &dto.RegisterResponse{
		Token: token,
	}, nil
}
