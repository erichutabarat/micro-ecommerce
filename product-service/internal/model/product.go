package model

import "gorm.io/gorm"

const (
	StatusDraft    = "draft"
	StatusActive   = "active"
	StatusArchived = "archived"
)

type Product struct {
	gorm.Model

	SellerID    uint   `gorm:"not null;index" json:"seller_id"`
	Name        string `gorm:"not null;size:200" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Price    int64  `gorm:"not null" json:"price"`
	Currency string `gorm:"type:char(3);not null;default:IDR" json:"currency"`

	Stock  int    `gorm:"not null;default:0" json:"stock"`
	Status string `gorm:"type:varchar(20);not null;default:draft;index" json:"status"`
}
