package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// ---------- Register ----------

type RegisterArgs struct {
	Email    string
	Password string
	Name     string
}

type RegisterReply struct {
	Success bool
	Message string
	User    *model.User
}

func (s *AuthService) Register(args *RegisterArgs, reply *RegisterReply) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		reply.Success = false
		reply.Message = "failed to process password"
		return nil
	}

	user := &model.User{
		Email:    args.Email,
		Password: string(hashed),
		Name:     args.Name,
	}

	result, err := s.repo.Register(user)
	if err != nil {
		reply.Success = false
		reply.Message = err.Error()
		return nil
	}

	result.Password = ""
	reply.Success = true
	reply.Message = "User registered successfully"
	reply.User = result
	return nil
}

// ---------- Login ----------

type LoginArgs struct {
	Email    string
	Password string
}

type LoginReply struct {
	Success bool
	Message string
	User    *model.User
}

func (s *AuthService) Login(args *LoginArgs, reply *LoginReply) error {
	user, err := s.repo.FindByEmail(args.Email)
	if err != nil {
		reply.Success = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reply.Message = "invalid email or password"
		} else {
			reply.Message = "internal error"
		}
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(args.Password)); err != nil {
		reply.Success = false
		reply.Message = "invalid email or password"
		return nil
	}

	user.Password = ""
	reply.Success = true
	reply.Message = "Login successful"
	reply.User = user
	return nil
}
