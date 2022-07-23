.PHONY: server
server:
	go build -o bin/api/server -v ./cmd/server

.PHONY: swagger
swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models