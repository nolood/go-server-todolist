FROM golang:1.21

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

EXPOSE 5000

CMD ["go", "run", "main.go"]