FROM golang:latest as go_dev

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main ./src/core/cmd/main.go

EXPOSE 8080

CMD ["./main"]