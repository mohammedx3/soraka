FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY cmd/server/main.go ./cmd/server/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
