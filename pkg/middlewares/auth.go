package middlewares

import (
	"errors"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/controllers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"github.com/juanvillacortac/entrenamiento-go/pkg/handlers"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "id"

func AuthMiddleware(r *gin.Engine) gin.HandlerFunc {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "songs api",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entities.UserSession); ok {
				return jwt.MapClaims{
					identityKey: v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &entities.UserSession{
				UserId: claims[identityKey].(uint),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var data entities.UserLogin
			if err := c.ShouldBind(&data); err != nil {
				return nil, errors.New("missing Email or Password")
			}

			user := controllers.GetUserByEmail(data.Email)

			if user != nil {
				err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.PasswordHash))
				if err != nil {
					return &entities.UserSession{
						UserId: user.ID,
					}, nil
				}
			}

			return nil, errors.New("incorrect Email or Password")
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*entities.UserSession); ok {
				var count int64
				db.DB.Model(&entities.User{}).Where("id = ?", v.UserId).Count(&count)
				return count > 0
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := middleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	auth := r.Group("/auth")
	{
		auth.GET("/refresh_token", middleware.RefreshHandler)
		auth.POST("/login", middleware.LoginHandler)
		auth.POST("/register", handlers.RegisterUserHandler)
	}

	return middleware.MiddlewareFunc()
}
