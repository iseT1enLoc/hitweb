package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/api/handler"
	"go_practice.com/component/appconfig"
	"go_practice.com/repository"
	"go_practice.com/usecase"
)

func NewSignInRoute(env *appconfig.Env, timeout time.Duration, db *sql.DB, r *gin.RouterGroup) {
	userrepo := repository.NewUserRepository(db, timeout)
	sign_handler := handler.SignInHandler{
		SignInUseCase: usecase.NewSignInUsecase(userrepo, timeout),
		Time:          timeout,
	}
	r.POST("/signin", sign_handler.SignIn())
}
