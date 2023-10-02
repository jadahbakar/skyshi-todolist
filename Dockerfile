# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.io
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY config.env .
COPY script/start.sh .
COPY script/wait-for.sh .
COPY log ./log
COPY db/migration ./migration


EXPOSE 8000
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]