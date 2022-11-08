package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string
	PasswordHash string
}

type UserSession struct {
	UserId uint
}

type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
