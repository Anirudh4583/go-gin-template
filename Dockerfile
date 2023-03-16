# Stage 1: Build the Go code
FROM golang:alpine as builder
# Install git.
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-template .

# Stage 2:
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the binary from the builder
COPY --from=builder /app/go-gin-template .
COPY --from=builder /app/.env .       
EXPOSE ${SERVER_HTTP_PORT} 
CMD ["./go-gin-template"]