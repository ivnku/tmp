FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./bin/site ./cmd/site/main.go

CMD ["./bin/site"]