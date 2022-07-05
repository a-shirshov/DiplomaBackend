FROM golang:1.16-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main cmd/base/main.go

FROM alpine:3
COPY --from=builder main /bin/main
EXPOSE 8080
ENTRYPOINT ["/bin/main"]