package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"go_practice.com/api/middleware"
	"go_practice.com/api/route"
	"go_practice.com/component/appconfig"
	"go_practice.com/component/appcontext"
	"go_practice.com/infras/postgres"
)

//go:embed
var embedMigrations embed.FS

func main() {
	fmt.Print("Hello world\n")
	appcfg, err := appconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading appconfig...[error]: %v", err)
		return
	}
	db, err := postgres.ConnectToDatabasein20s(appcfg)
	if err != nil {
		log.Fatalf("Fail to connect to database after 20s, [error]: %v", err)
	}
	cfg := appcontext.NewAppContext(db, os.Getenv("SECRET_KEY"))

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.Recover(cfg))
	route.SetUp(appcfg, 3600, db, r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Successfully set up"})
	})

	r.Run("localhost:8080")
}
