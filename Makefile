
up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} up
down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} down
status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${dsn} goose -dir=${migrationDir} status
run:
	go run cmd/main.go