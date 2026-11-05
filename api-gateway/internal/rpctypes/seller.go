package rpctypes

type Seller struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	StoreName   string `json:"store_name"`
	Description string `json:"description"`
}

type SellerRegisterArgs struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	StoreName   string `json:"store_name"`
	Description string `json:"description"`
}

type SellerRegisterReply struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Seller  *Seller `json:"seller"`
}

type SellerLoginArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerLoginReply struct {
	Success   bool    `json:"success"`
	Message   string  `json:"message"`
	Seller    *Seller `json:"seller"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expires_at"`
}

func (r *SellerRegisterReply) IsSuccess() bool    { return r.Success }
func (r *SellerRegisterReply) GetMessage() string { return r.Message }
func (r *SellerLoginReply) IsSuccess() bool       { return r.Success }
func (r *SellerLoginReply) GetMessage() string    { return r.Message }
