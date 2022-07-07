FROM golang:1.16-alpine as main_server_build
RUN apk add --no-cache make 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make main_server
EXPOSE 8080
WORKDIR /app/bin/api
CMD ["./main_server"]