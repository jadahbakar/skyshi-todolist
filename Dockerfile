# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.io
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY config.env .
COPY log ./log
COPY db/migrations ./db/migrations


EXPOSE 3030
CMD [ "/app/main" ]
