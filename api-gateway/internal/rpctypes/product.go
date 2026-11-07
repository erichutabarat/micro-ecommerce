package rpctypes

import "time"

type Product struct {
	ID          uint      `json:"id"`
	SellerID    uint      `json:"seller_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Currency    string    `json:"currency"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SellerScoped interface {
	SetSellerID(id uint)
}

type CreateProductArgs struct {
	SellerID    uint   `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Currency    string `json:"currency"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
}

type GetProductArgs struct {
	ID uint `json:"id"`
}

type ListProductsArgs struct {
	Search   string `json:"search"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type ListMyProductsArgs struct {
	SellerID uint   `json:"-"`
	Status   string `json:"status"`
	Search   string `json:"search"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type UpdateProductArgs struct {
	SellerID    uint   `json:"-"`
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Currency    string `json:"currency"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
}

type DeleteProductArgs struct {
	SellerID uint `json:"-"`
	ID       uint `json:"id"`
}

func (a *CreateProductArgs) SetSellerID(id uint)  { a.SellerID = id }
func (a *ListMyProductsArgs) SetSellerID(id uint) { a.SellerID = id }
func (a *UpdateProductArgs) SetSellerID(id uint)  { a.SellerID = id }
func (a *DeleteProductArgs) SetSellerID(id uint)  { a.SellerID = id }

type ProductReply struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Product *Product `json:"product"`
}

type DeleteProductReply struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ListProductsReply struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message"`
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

func (r *ProductReply) IsSuccess() bool          { return r.Success }
func (r *ProductReply) GetMessage() string       { return r.Message }
func (r *DeleteProductReply) IsSuccess() bool    { return r.Success }
func (r *DeleteProductReply) GetMessage() string { return r.Message }
func (r *ListProductsReply) IsSuccess() bool     { return r.Success }
func (r *ListProductsReply) GetMessage() string  { return r.Message }
