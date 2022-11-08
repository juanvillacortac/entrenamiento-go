package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/controllers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(c *gin.Context) {
	var data entities.UserLogin
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := controllers.GetUserByEmail(data.Email)

	if user != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "email taken",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&entities.User{
		Email:        data.Email,
		PasswordHash: string(hash),
	})
	c.JSON(http.StatusOK, gin.H{
		"status": "registered",
	})
}
