package appcontext

import (
	"database/sql"
)

type AppContext interface {
	GetConnectionToDatabase() *sql.DB
	GetSecretKeyString() string
}

type appcontext struct {
	db        *sql.DB
	secretkey string
}

func NewAppContext(db *sql.DB, secretkey string) *appcontext {
	return &appcontext{
		db:        db,
		secretkey: secretkey,
	}
}
func (appctx *appcontext) GetConnectionToDatabase() *sql.DB {
	return appctx.db
}
func (appctx *appcontext) GetSecretKeyString() string {
	return appctx.secretkey
}
