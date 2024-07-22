include config.env

DSN:="host=$(PG_HOST) port=$(PG_PORT) user=$(PG_USER) password=$(PG_PWD) dbname=$(PG_DBNAME) sslmode=disable"
	
run:
	goose -allow-missing -dir ./database/migrations postgres $(DSN) up
	cd ./cmd/time-tracker && go run main.go

	
	