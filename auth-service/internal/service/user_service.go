package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/internal/token"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo   repository.UserRepository
	tokens *token.Manager
}

func NewAuthService(repo repository.UserRepository, tokens *token.Manager) *AuthService {
	return &AuthService{repo: repo, tokens: tokens}
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
	Success   bool
	Message   string
	User      *model.User
	Token     string
	ExpiresAt int64
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

	signed, expiresAt, err := s.tokens.Generate(user.ID, user.Email)
	if err != nil {
		log.Printf("failed to generate token: %v", err)
		reply.Success = false
		reply.Message = "internal error"
		return nil
	}

	// log.Printf("User %s logged in successfully, token expires at %v", user.Email, expiresAt)

	user.Password = ""
	reply.Success = true
	reply.Message = "Login successful"
	reply.User = user
	reply.Token = signed
	reply.ExpiresAt = expiresAt.Unix()
	return nil
}
