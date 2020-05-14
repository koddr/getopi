.PHONY: back-start

back-start: build-api-server test-api-server
	./build/api_server

build-api-server:
	@rm -rf ./build
	@GOARCH=amd64 CGOENABLED=0 go build -ldflags="-w -s" -o ./build/api_server ./main.go

test-api-server:
	godotenv -f .env go test -v ./... -cover

migrate-up:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" down

migrate-force:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" force $(v)