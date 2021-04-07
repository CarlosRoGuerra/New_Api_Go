FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o bin/main cmd/main.go

EXPOSE 8000

CMD ["/app/bin/main"]