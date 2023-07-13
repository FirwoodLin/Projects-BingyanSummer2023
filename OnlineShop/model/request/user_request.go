package request

// UserRequest 注册时的请求
type UserRequest struct {
	Name     string `validate:"required,min=3,max=20" json:"name"`
	Password string `validate:"required,min=8,max=20" json:"password"`
	Email    string `validate:"required,email" json:"email"`
	Tel      string `validate:"required,e164" json:"tel"` // E.164 标准:国际关于手机号的规范
	Nickname string `validate:"required,max=20" json:"nickname"`
	UserID   uint   `json:"-"`
	IsAdmin  bool   `json:"-"`
}

// UserSignInRequest 登陆时的请求
type UserSignInRequest struct {
	// UserID       string `json:"-"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=20" json:"password"`
}

type UserUpdateRequest struct {
	Name     string `validate:"omitempty,min=3,max=20" json:"name,omitempty"`
	Password string `validate:"omitempty,min=8,max=20" json:"password,omitempty"`
	Email    string `validate:"omitempty,email" json:"email,omitempty"`
	Tel      string `validate:"omitempty,e164" json:"tel,omitempty"` // E.164 标准:国际关于手机号的规范
	Nickname string `validate:"omitempty,max=20" json:"nickname,omitempty"`
	UserID   uint   `json:"-"`
}
