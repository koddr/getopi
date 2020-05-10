.PHONY: back-start

back-start:
	go run main.go

front-start:
	@cd ./frontend && npm start
	@echo "[OK] Preact app is running!

migrate-up:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" down

migrate-force:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" force $(v)

tests:
	godotenv -f .env go test -v ./... -cover