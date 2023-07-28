FROM mysql:8.0.32-debian

RUN apt-get update \
    && apt-get install -y locales \
    && sed -i -E 's/# (ja_JP.UTF-8)/\1/' /etc/locale.gen \
    && locale-gen \
    && update-locale LANG=ja_JP.UTF-8 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
ENV LC_ALL ja_JP.UTF-8


# ------------------
# Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .
# # COPY src/core/cmd* ./
# COPY src/core/infra/controllers* ./


# RUN go mod init example.com/myapp
RUN go mod download
RUN go get github.com/gin-gonic/gin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/core/cmd/main.go

EXPOSE 8080

CMD ["./main"]

