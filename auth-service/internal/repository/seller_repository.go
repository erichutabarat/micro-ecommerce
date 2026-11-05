package repository

import (
	"auth-service/internal/model"

	"gorm.io/gorm"
)

type SellerRepository interface {
	Register(seller *model.Seller) (*model.Seller, error)
	FindByEmail(email string) (*model.Seller, error)
}

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) SellerRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) Register(seller *model.Seller) (*model.Seller, error) {
	if err := r.db.Create(seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (r *sellerRepository) FindByEmail(email string) (*model.Seller, error) {
	var seller model.Seller
	if err := r.db.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}
