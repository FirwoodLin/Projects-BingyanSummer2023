package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client
var RedisCtx context.Context

type User struct {
	Name           string `gorm:"type:varchar(20);unique"`
	Nickname       string `gorm:"type:varchar(20)"`
	Email          string `gorm:"unique"`
	Tel            string `gorm:"unique"`
	HashedPassword string `gorm:"size:60"`
	IsAdmin        bool
	// ID             uint `gorm:"primarykey"`
	gorm.Model
}
type UserRequset struct {
	Name     string `validate:"required,min=3,max=20" json:"name"`
	Password string `validate:"required,min=8,max=20" json:"password"`
	Email    string `validate:"required,email" json:"email"`
	Tel      string `validate:"required,e164" json:"tel"` // E.164 标准:国际关于手机号的规范
	Nickname string `validate:"required,max=20" json:"nickname"`
	ID       uint   `json:"-"`
	IsAdmin  bool   `json:"-"`
}
type UserUpdateRequset struct {
	Name     string `validate:"omitempty,min=3,max=20" json:"name,omitempty"`
	Password string `validate:"omitempty,min=8,max=20" json:"password,omitempty"`
	Email    string `validate:"omitempty,email" json:"email,omitempty"`
	Tel      string `validate:"omitempty,e164" json:"tel,omitempty"` // E.164 标准:国际关于手机号的规范
	Nickname string `validate:"omitempty,max=20" json:"nickname,omitempty"`
	ID       uint   `json:"-"`
}
type UserSignInRequest struct {
	// ID       string `json:"-"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=20" json:"password"`
}
type UserSignInResponse struct {
	ID      uint
	IsAdmin bool
}

type UserQueryResponse struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	IsAdmin  bool   `json:"is_admin"`
	ID       uint   `json:"id"`
}

func ConvertUserRequestToUser(req *UserRequset) *User {
	user := &User{}
	user.Name = req.Name
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.Tel = req.Tel
	return user
}

// func ConvertUserUpdateRequsetToUser(req *UserUpdateRequset) *User {
// 	user := &User{}
// 	user.Name = req.Name
// 	user.Nickname = req.Nickname
// 	user.Email = req.Email
// 	user.Tel = req.Tel
// 	user.HashedPassword =
// 	return user
// }
