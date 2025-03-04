# The build stage
FROM golang:1.24-alpine AS builder
WORKDIR /quota-ms/app/
COPY /app/quota.go ./
COPY /app/go.mod ./
COPY /app/go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o go-api quota.go
RUN go build -o go-api quota.go

# The run stage
FROM alpine:3.21
WORKDIR /quota-ms
COPY --from=builder /quota-ms/app .
EXPOSE 3000:8080
CMD ["./go-api"]