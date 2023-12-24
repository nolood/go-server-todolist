FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]