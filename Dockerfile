# build stage
FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rental-service cmd/rent/main.go
RUN ls /app

# runtime stage
FROM alpine:latest

WORKDIR /root
COPY --from=builder /app/rental-service .

EXPOSE 8080
CMD [ "./rental-service" ]