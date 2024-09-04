package appconfig

type Env struct {
	DB_HOST             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	PORT                string
	SECRET_KEY          string
	SECRET_ACCESSKEY    string
	SECRET_REFRESHKEY   string
	EXPIRATION_TIME_SEC string
}
