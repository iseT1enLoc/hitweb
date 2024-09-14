package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"go_practice.com/component/appconfig"
)

func ConnectToDatabasein20s(appcfg *appconfig.Env) (*sql.DB, error) {
	timeTry := time.Second * 20
	//os.Getenv("GOOSE_DBSTRING")
	//create a connection function
	connectingDatabase := func(appcfig *appconfig.Env) (*sql.DB, error) {
		//dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", appcfg.DB_HOST, appcfg.DB_USER, appcfg.DB_PASSWORD, appcfg.DB_NAME, appcfg.PORT)
		psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
			"password=%v dbname=%v sslmode=disable",
			appcfg.DB_HOST, appcfg.PORT, appcfg.DB_USER, appcfg.DB_PASSWORD, appcfg.DB_NAME)
		//dsn := os.Getenv("remotedbString")
		//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			//panic(err)
			fmt.Printf("Error happened while connecting to database [error]-%v", err)
			return nil, err
		}
		//return db, nil
		//defer db.Close()

		// Test the connection with db.Ping()
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		fmt.Print("Connect successfully to database")
		return db, nil
	}
	deadline := time.Now().Add(timeTry)
	var db *sql.DB
	var err error

	for time.Now().Before(deadline) {
		log.Println("CONNECT to database.....")
		db, err = connectingDatabase(appcfg)
		if err == nil {
			//fmt.Printf("Database name: %v", db.Name())
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("Error while connecting to database...[error]: %v", err)
}
