FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o mm-rest-api .

FROM postgres:latest

WORKDIR /usr/local/bin

COPY --from=builder /app/mm-rest-api .
COPY --from=builder /app/.env .

CMD ["./mm-rest-api"]
