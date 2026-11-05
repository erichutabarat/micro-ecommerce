package model

import "gorm.io/gorm"

type Seller struct {
	gorm.Model

	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"-"`
	Name        string `gorm:"not null" json:"name"`
	StoreName   string `gorm:"unique;not null" json:"store_name"`
	Description string `json:"description"`
}
