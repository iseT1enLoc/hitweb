package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/component/appconfig"
	"gorm.io/gorm"
)

func SetUp(env *appconfig.Env, timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	publicRoute := r.Group("api/public/")
	NewSignUpRoute(env, timeout, db, publicRoute)
}
