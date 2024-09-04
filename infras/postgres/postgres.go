package postgres

import (
	"fmt"
	"log"
	"time"

	"go_practice.com/component/appconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabasein20s(appcfg *appconfig.Env) (*gorm.DB, error) {
	timeTry := time.Second * 20
	//create a connection function
	connectingDatabase := func(appcfig *appconfig.Env) (*gorm.DB, error) {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", appcfg.DB_HOST, appcfg.DB_USER, appcfg.DB_PASSWORD, appcfg.DB_NAME, appcfg.PORT)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db.Debug(), nil
	}
	deadline := time.Now().Add(timeTry)
	var db *gorm.DB
	var err error
	for time.Now().Before(deadline) {
		log.Println("CONNECT to database.....")
		db, err = connectingDatabase(appcfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("Error while connecting to database...[error]: %v", err)
}
