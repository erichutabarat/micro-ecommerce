package rpctypes

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type RegisterArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RegisterReply struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type LoginArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReply struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	User      *User  `json:"user"`
	Token     string
	ExpiresAt int64
}

// Lets the handler read any reply generically.
func (r *RegisterReply) IsSuccess() bool    { return r.Success }
func (r *RegisterReply) GetMessage() string { return r.Message }
func (r *LoginReply) IsSuccess() bool       { return r.Success }
func (r *LoginReply) GetMessage() string    { return r.Message }

type Result interface {
	IsSuccess() bool
	GetMessage() string
}
