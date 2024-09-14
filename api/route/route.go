package route

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/api/middleware"
	"go_practice.com/component/appconfig"
	"gorm.io/gorm"
)

func SetUp(env *appconfig.Env, timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	publicRoute := r.Group("api/public/")
	protectedRoute := r.Group("api/protected/")
	NewSignUpRoute(env, timeout, db, publicRoute)
	NewSignInRoute(env, timeout, db, publicRoute)
	protectedRoute.Use(middleware.JwtAuthMiddleware(os.Getenv("SECRET_KEY")))
	NewResourceRoute(env, timeout, db, protectedRoute)
}
