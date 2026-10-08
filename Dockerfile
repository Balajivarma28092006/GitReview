FROM golang:1.27-alpine AS builder

RUN apk add --no-cache git CA-certificates

WORKDIR /app 

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go 

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /root/ 

COPY --from=builder /app/server .
EXPOSE 8080

CMD ["./server"]
