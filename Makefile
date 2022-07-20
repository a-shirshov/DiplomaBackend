.PHONY: server
server:
	go build -o bin/api/server -v ./cmd/server