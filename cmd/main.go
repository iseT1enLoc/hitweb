package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go_practice.com/component/appconfig"
	"go_practice.com/component/appcontext"
	"go_practice.com/infras/postgres"
)

func main() {
	fmt.Print("Hello world\n")
	appcfg, err := appconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading appconfig...[error]: %v", err)
		return
	}
	db, err := postgres.ConnectToDatabasein20s(appcfg)
	_ = appcontext.NewAppContext(db, os.Getenv("SECRET_KEY"))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Successfully set up"})
	})

	r.Run("localhost:8080")
}
