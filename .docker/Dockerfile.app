FROM golang:latest as go_dev

WORKDIR /app

COPY . .

RUN go mod download
RUN go get github.com/gin-gonic/gin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/core/cmd/main.go

EXPOSE 8080

CMD ["./main"]