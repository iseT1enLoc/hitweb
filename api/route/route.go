package route

import (
	"database/sql"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/api/middleware"
	"go_practice.com/component/appconfig"
)

func SetUp(env *appconfig.Env, timeout time.Duration, db *sql.DB, r *gin.Engine) {
	publicRoute := r.Group("api/public/")
	protectedRoute := r.Group("api/protected/")
	NewSignUpRoute(env, timeout, db, publicRoute)
	NewSignInRoute(env, timeout, db, publicRoute)
	protectedRoute.Use(middleware.JwtAuthMiddleware(os.Getenv("SECRET_KEY")))
	NewResourceRoute(env, timeout, db, protectedRoute)
}
