# The build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o go-api /app/quota.go

# The run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/quota .
EXPOSE 3000
CMD ["./quota"]