FROM golang:1.22

WORKDIR /auth_api

COPY . .

RUN go mod tidy

ENTRYPOINT [ "go run ./cmd/auth_api.go" ]