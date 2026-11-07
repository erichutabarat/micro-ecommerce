package service

import (
	"errors"
	"strings"

	"product-service/internal/model"
	"product-service/internal/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// ---------- shared reply types ----------

type ProductReply struct {
	Success bool
	Message string
	Product *model.Product
}

type DeleteProductReply struct {
	Success bool
	Message string
}

type ListProductsReply struct {
	Success  bool
	Message  string
	Products []model.Product
	Total    int64
	Page     int
	PageSize int
}

// ---------- helpers ----------

const (
	msgUnauthorized = "unauthorized"
	msgNotFound     = "product not found"
	msgInternal     = "internal error"
)

func validStatus(s string) bool {
	return s == model.StatusDraft || s == model.StatusActive || s == model.StatusArchived
}

// validate returns an error message, or "" if the input is fine.
func validate(name string, price int64, currency string, stock int, status string) string {
	switch {
	case strings.TrimSpace(name) == "":
		return "name is required"
	case len(name) > 200:
		return "name must be at most 200 characters"
	case price < 0:
		return "price must not be negative"
	case len(currency) != 3:
		return "currency must be a 3-letter code"
	case stock < 0:
		return "stock must not be negative"
	case !validStatus(status):
		return "status must be draft, active or archived"
	}
	return ""
}

func paginate(page, size int) (offset, limit, p, s int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	if size > 100 {
		size = 100
	}
	return (page - 1) * size, size, page, size
}

func normalize(currency, status string) (string, string) {
	currency = strings.ToUpper(strings.TrimSpace(currency))
	if currency == "" {
		currency = "IDR"
	}
	if status == "" {
		status = model.StatusDraft
	}
	return currency, status
}

// ---------- Create (seller) ----------

type CreateProductArgs struct {
	SellerID    uint // set by the gateway from the JWT, never from the client
	Name        string
	Description string
	Price       int64
	Currency    string
	Stock       int
	Status      string
}

func (s *ProductService) Create(args *CreateProductArgs, reply *ProductReply) error {
	if args.SellerID == 0 {
		reply.Message = msgUnauthorized
		return nil
	}

	currency, status := normalize(args.Currency, args.Status)
	if msg := validate(args.Name, args.Price, currency, args.Stock, status); msg != "" {
		reply.Message = msg
		return nil
	}

	result, err := s.repo.Create(&model.Product{
		SellerID:    args.SellerID,
		Name:        strings.TrimSpace(args.Name),
		Description: args.Description,
		Price:       args.Price,
		Currency:    currency,
		Stock:       args.Stock,
		Status:      status,
	})
	if err != nil {
		reply.Message = msgInternal
		return nil
	}

	reply.Success = true
	reply.Message = "Product created successfully"
	reply.Product = result
	return nil
}

// ---------- Get (public: active products only) ----------

type GetProductArgs struct {
	ID uint
}

func (s *ProductService) Get(args *GetProductArgs, reply *ProductReply) error {
	p, err := s.repo.FindByID(args.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reply.Message = msgNotFound
		} else {
			reply.Message = msgInternal
		}
		return nil
	}

	// drafts and archived products are invisible to the public
	if p.Status != model.StatusActive {
		reply.Message = msgNotFound
		return nil
	}

	reply.Success = true
	reply.Message = "OK"
	reply.Product = p
	return nil
}

// ---------- List (public: active products only) ----------

type ListProductsArgs struct {
	Search   string
	Page     int
	PageSize int
}

func (s *ProductService) List(args *ListProductsArgs, reply *ListProductsReply) error {
	offset, limit, page, size := paginate(args.Page, args.PageSize)

	items, total, err := s.repo.List(repository.ListFilter{
		Status: model.StatusActive,
		Search: args.Search,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		reply.Message = msgInternal
		return nil
	}

	reply.Success = true
	reply.Message = "OK"
	reply.Products = items
	reply.Total = total
	reply.Page = page
	reply.PageSize = size
	return nil
}

// ---------- ListMine (seller: all statuses, own products) ----------

type ListMyProductsArgs struct {
	SellerID uint
	Status   string // optional filter
	Search   string
	Page     int
	PageSize int
}

func (s *ProductService) ListMine(args *ListMyProductsArgs, reply *ListProductsReply) error {
	if args.SellerID == 0 {
		reply.Message = msgUnauthorized
		return nil
	}
	if args.Status != "" && !validStatus(args.Status) {
		reply.Message = "status must be draft, active or archived"
		return nil
	}

	offset, limit, page, size := paginate(args.Page, args.PageSize)

	items, total, err := s.repo.List(repository.ListFilter{
		SellerID: args.SellerID,
		Status:   args.Status,
		Search:   args.Search,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		reply.Message = msgInternal
		return nil
	}

	reply.Success = true
	reply.Message = "OK"
	reply.Products = items
	reply.Total = total
	reply.Page = page
	reply.PageSize = size
	return nil
}

// ---------- Update (seller, own products only) ----------

// Full replacement: send every field. See the note on gob below.
type UpdateProductArgs struct {
	SellerID    uint
	ID          uint
	Name        string
	Description string
	Price       int64
	Currency    string
	Stock       int
	Status      string
}

func (s *ProductService) Update(args *UpdateProductArgs, reply *ProductReply) error {
	if args.SellerID == 0 {
		reply.Message = msgUnauthorized
		return nil
	}

	currency, status := normalize(args.Currency, args.Status)
	if msg := validate(args.Name, args.Price, currency, args.Stock, status); msg != "" {
		reply.Message = msg
		return nil
	}

	p, err := s.repo.FindOwned(args.ID, args.SellerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reply.Message = msgNotFound
		} else {
			reply.Message = msgInternal
		}
		return nil
	}

	p.Name = strings.TrimSpace(args.Name)
	p.Description = args.Description
	p.Price = args.Price
	p.Currency = currency
	p.Stock = args.Stock
	p.Status = status

	result, err := s.repo.Save(p)
	if err != nil {
		reply.Message = msgInternal
		return nil
	}

	reply.Success = true
	reply.Message = "Product updated successfully"
	reply.Product = result
	return nil
}

// ---------- Delete (seller, own products only) ----------

type DeleteProductArgs struct {
	SellerID uint
	ID       uint
}

func (s *ProductService) Delete(args *DeleteProductArgs, reply *DeleteProductReply) error {
	if args.SellerID == 0 {
		reply.Message = msgUnauthorized
		return nil
	}

	if err := s.repo.Delete(args.ID, args.SellerID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reply.Message = msgNotFound
		} else {
			reply.Message = msgInternal
		}
		return nil
	}

	reply.Success = true
	reply.Message = "Product deleted successfully"
	return nil
}
