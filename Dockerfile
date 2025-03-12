#------ Step 1: Build Stage --------
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o client-go .

# -------Step 2: Run Stage -----------
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/client-go .
CMD ["./client-go"]


