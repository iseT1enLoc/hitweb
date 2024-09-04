dsn="host=localhost user=postgres password=0123456789 dbname=website port=5432 sslmode=disable TimeZone=Asia/Shanghai"
migrationDir="db/migrations/"
up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} up
down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} down
status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} status
run:
	go run cmd/main.go