############################
# STEP 1: Build executable binary
############################
FROM golang:1.21-alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o myapp .

############################
# STEP 2: Build a small image
############################
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/myapp /app/myapp

EXPOSE 8080

CMD ["./myapp"]
