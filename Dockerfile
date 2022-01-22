FROM golang:latest as builder

COPY . /github.com/ievseev/fiboslicer/
WORKDIR /github.com/ievseev/fiboslicer/

RUN go mod download && GOOS=linux go build -o ./bin/fiboslicer ./cmd/main.go

EXPOSE 80

CMD ["./bin/fiboslicer"]

