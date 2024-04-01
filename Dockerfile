FROM golang:1.22

WORKDIR /auth_api

COPY . .

RUN go mod tidy

CMD ["go", "run", "./cmd/auth_api.go"]