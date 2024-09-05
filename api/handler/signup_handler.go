package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/domain"
)

type SignUpHandler struct {
	SignUpusecase domain.ISignUpUseCase
	Timeout       time.Duration
}

func (sh *SignUpHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signupReq domain.SignUpReq
		err := c.ShouldBind(&signupReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		err = sh.SignUpusecase.GetUserByEmail(signupReq.UserEmail)
		if err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Email has already existed in DB"})
			return
		}
		user, err := sh.SignUpusecase.SignUp(signupReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		accesstoken, err := sh.SignUpusecase.CreateAcessToken(1, os.Getenv("SECRET_KEY"), user)
		if err != nil {
			log.Fatalf("Fail to create access token at user handler [error]-%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		refreshtoken, err := sh.SignUpusecase.CreateRefreshToken(1, os.Getenv("SECRET_KEY"), user)
		if err != nil {
			log.Fatalf("fail tp create refresh token [error]-%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"access_token": accesstoken, "refresh_token": refreshtoken})
	}
}
