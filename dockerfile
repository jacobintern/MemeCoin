FROM golang:1.22.5

# default at /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]