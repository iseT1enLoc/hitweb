package appcontext

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetConnectionToDatabase() *gorm.DB
	GetSecretKeyString() string
}

type appcontext struct {
	db        *gorm.DB
	secretkey string
}

func NewAppContext(db *gorm.DB, secretkey string) *appcontext {
	return &appcontext{
		db:        db,
		secretkey: secretkey,
	}
}
func (appctx *appcontext) GetConnectionToDatabase() *gorm.DB {
	return appctx.db
}
func (appctx *appcontext) GetSecretKeyString() string {
	return appctx.secretkey
}
