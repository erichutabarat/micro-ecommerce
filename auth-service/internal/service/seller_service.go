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

type SellerService struct {
	repo   repository.SellerRepository
	tokens *token.Manager
}

func NewSellerService(repo repository.SellerRepository, tokens *token.Manager) *SellerService {
	return &SellerService{repo: repo, tokens: tokens}
}

// ---------- Register ----------

type SellerRegisterArgs struct {
	Email       string
	Password    string
	Name        string
	StoreName   string
	Description string
}

type SellerRegisterReply struct {
	Success bool
	Message string
	Seller  *model.Seller
}

func (s *SellerService) Register(args *SellerRegisterArgs, reply *SellerRegisterReply) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		reply.Message = "failed to process password"
		return nil
	}

	seller := &model.Seller{
		Email:       args.Email,
		Password:    string(hashed),
		Name:        args.Name,
		StoreName:   args.StoreName,
		Description: args.Description,
	}

	result, err := s.repo.Register(seller)
	if err != nil {
		reply.Message = err.Error()
		return nil
	}

	result.Password = ""
	reply.Success = true
	reply.Message = "Seller registered successfully"
	reply.Seller = result
	return nil
}

// ---------- Login ----------

type SellerLoginArgs struct {
	Email    string
	Password string
}

type SellerLoginReply struct {
	Success   bool
	Message   string
	Seller    *model.Seller
	Token     string
	ExpiresAt int64
}

func (s *SellerService) Login(args *SellerLoginArgs, reply *SellerLoginReply) error {
	seller, err := s.repo.FindByEmail(args.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reply.Message = "invalid email or password"
		} else {
			reply.Message = "internal error"
		}
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(args.Password)); err != nil {
		reply.Message = "invalid email or password"
		return nil
	}

	signed, expiresAt, err := s.tokens.Generate(seller.ID, seller.Email, token.RoleSeller)
	if err != nil {
		log.Printf("failed to generate seller token: %v", err)
		reply.Message = "internal error"
		return nil
	}

	seller.Password = ""
	reply.Success = true
	reply.Message = "Login successful"
	reply.Seller = seller
	reply.Token = signed
	reply.ExpiresAt = expiresAt.Unix()
	return nil
}
