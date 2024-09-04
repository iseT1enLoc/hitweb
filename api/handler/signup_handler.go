package handler

import (
	"net/http"
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

		username, err := sh.SignUpusecase.SignUp(signupReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": username})
	}
}
