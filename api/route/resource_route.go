package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/component/appconfig"
	"gorm.io/gorm"
)

func NewResourceRoute(env *appconfig.Env, timeout time.Duration, db *gorm.DB, r *gin.RouterGroup) {
	r.GET("/resource", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "secure resource..."})
	})
}
