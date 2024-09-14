package route

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go_practice.com/component/appconfig"
)

func NewResourceRoute(env *appconfig.Env, timeout time.Duration, db *sql.DB, r *gin.RouterGroup) {
	r.GET("/resource", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "secure resource..."})
	})
}
