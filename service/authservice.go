package service

import (
	"mygram/entity"
	userrepository "mygram/repository/UserRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo userrepository.UserRepository
}

func NewAuthService(userRepo userrepository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		var user *entity.User = &entity.User{}

		tokenStr := ctx.Request.Header.Get("Authorization")

		err := user.VerifyToken(tokenStr)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": err.Error(),
			})
			return
		}

		_, err = a.userRepo.FindByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": err.Error(),
			})
			return
		}

		ctx.Set("email", &user.Email)
		ctx.Next()
	})
}
