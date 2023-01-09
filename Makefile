.PHONY: server
server:
	go build -tags=jsoniter -o bin/api/server -v ./cmd/server 

.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go

.PHONY: test
test:
	go test -race ./... 

.PHONY: cover
cover:
	go test -cover -coverprofile=cover.out -coverpkg=./... ./...
	cat cover.out | fgrep -v "main.go" | fgrep -v "mock.go" | fgrep -v "docs.go" | fgrep -v "logger.go"  > cover1.out
	go tool cover -func=cover1.out