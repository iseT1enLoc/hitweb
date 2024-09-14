package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwtutils "go_practice.com/component/jwt_utils"
	"go_practice.com/domain"
)

type SignInHandler struct {
	SignInUseCase domain.ISignIn
	Time          time.Duration
}

func (s *SignInHandler) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signInReq domain.SignInReq

		err := c.ShouldBind(&signInReq)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "can not bind json"})
			return
		}
		desisedUser, err := s.SignInUseCase.GetUserByEmail(signInReq.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "The requested Email do not exist in database"})
			return
		}
		ok := jwtutils.CheckPassword(desisedUser.Password, signInReq.Password)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "The requested password do not valid"})
			return
		}
		accesstoken, err := jwtutils.CreateAcessToken(1, os.Getenv("SECRET_KEY"), desisedUser)
		if err != nil {
			log.Fatalf("Fail to create access token at user handler [error]-%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		refreshtoken, err := jwtutils.CreateRefreshToken(1, os.Getenv("SECRET_KEY"), desisedUser)
		if err != nil {
			log.Fatalf("fail tp create refresh token [error]-%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, domain.SignInResponse{
			AccessToken:  accesstoken,
			RefreshToken: refreshtoken,
		})
	}
}
