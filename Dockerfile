FROM golang:1.17.6-alpine3.15 as builder

COPY . /github.com/ievseev/fibonacci-slicer/
WORKDIR /github.com/ievseev/fibonacci-slicer/

RUN go mod download && GOOS=linux go build -o ./bin/fibonacci_slicer ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ievseev/fibonacci-slicer/bin/fibonacci_slicer .

EXPOSE 80

CMD ["./fibonacci_slicer"]

