package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/api/handler"
	"go_practice.com/component/appconfig"
	"go_practice.com/repository"
	"go_practice.com/usecase"
	"gorm.io/gorm"
)

func NewSignUpRoute(env *appconfig.Env, timeout time.Duration, db *gorm.DB, r *gin.RouterGroup) {
	repo := repository.NewUserRepository(db, timeout)
	sc := handler.SignUpHandler{
		SignUpusecase: usecase.NewSignUpUseCase(repo, timeout),
		Timeout:       timeout,
	}
	r.POST("/signup", sc.SignUp())
}
