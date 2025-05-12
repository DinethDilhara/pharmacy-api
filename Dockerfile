FROM golang:1.23.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@v1.40.4

COPY . .
RUN go mod tidy
