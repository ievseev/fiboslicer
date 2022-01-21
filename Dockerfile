FROM golang:latest as builder

COPY . /github.com/ievseev/fibonacci-slicer/
WORKDIR /github.com/ievseev/fibonacci-slicer/

RUN go mod download && GOOS=linux go build -o ./bin/fibonacci_slicer ./cmd/main.go

EXPOSE 80

CMD ["./bin/fibonacci_slicer"]

