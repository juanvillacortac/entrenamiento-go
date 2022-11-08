package controllers

import (
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
)

func GetUserByEmail(email string) *entities.User {
	var user entities.User

	db.DB.Where(&entities.User{
		Email: email,
	}).First(&user)

	if user.Email != "" {
		return &user
	}
	return nil
}
