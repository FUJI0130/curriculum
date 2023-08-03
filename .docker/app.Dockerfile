FROM golang:latest as go_dev

WORKDIR /app

# go.modに載るよう、先にgo get gin～　したので不要になった
# RUN go get github.com/gin-gonic/gin
COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main ./src/core/cmd/main.go

EXPOSE 8080

CMD ["./main"]