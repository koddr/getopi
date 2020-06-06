.PHONY: air-hot-reload

air-hot-reload:
	cd ./api && air -d

back-start: test-api-server build-api-server
	./api/build/api

front-start: 
	cd ./static && npm run dev

build-api-server:
	cd ./api \
	&& rm -rf ./build \
	&& GOARCH=amd64 CGOENABLED=0 go build -ldflags="-w -s" -o ./build/api .

test-api-server:
	cd ./api \
	&& godotenv -f ../.env go test -v ./... -cover

migrate-up:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" down

migrate-force:
	migrate -path migrations -database "postgres://koddr@localhost/koddr?sslmode=disable" force $(v)