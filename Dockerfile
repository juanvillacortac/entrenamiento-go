FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o ./bin/server ./cmd/server

CMD ./bin/server