# Development run
run-dev:
	@nodemon --exec go run main.go --signal SIGTERM

seed:
	@go run src/helpers/seeder/seeder.go

migrate:
	@go run src/helpers/migrate/migration.go
