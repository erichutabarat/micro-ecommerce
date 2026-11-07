package repository

import (
	"product-service/internal/model"

	"gorm.io/gorm"
)

type ListFilter struct {
	SellerID uint
	Status   string
	Search   string
	Offset   int
	Limit    int
}

type ProductRepository interface {
	Create(p *model.Product) (*model.Product, error)
	FindByID(id uint) (*model.Product, error)
	FindOwned(id, sellerID uint) (*model.Product, error)
	Save(p *model.Product) (*model.Product, error)
	Delete(id, sellerID uint) error
	List(f ListFilter) ([]model.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(p *model.Product) (*model.Product, error) {
	if err := r.db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var p model.Product
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) FindOwned(id, sellerID uint) (*model.Product, error) {
	var p model.Product
	if err := r.db.Where("id = ? AND seller_id = ?", id, sellerID).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Save(p *model.Product) (*model.Product, error) {
	if err := r.db.Save(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *productRepository) Delete(id, sellerID uint) error {
	res := r.db.Where("id = ? AND seller_id = ?", id, sellerID).Delete(&model.Product{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *productRepository) List(f ListFilter) ([]model.Product, int64, error) {
	q := r.db.Model(&model.Product{})
	if f.SellerID != 0 {
		q = q.Where("seller_id = ?", f.SellerID)
	}
	if f.Status != "" {
		q = q.Where("status = ?", f.Status)
	}
	if f.Search != "" {
		q = q.Where("name LIKE ?", "%"+f.Search+"%")
	}
	q = q.Session(&gorm.Session{})
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []model.Product
	if err := q.Order("id DESC").Offset(f.Offset).Limit(f.Limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
