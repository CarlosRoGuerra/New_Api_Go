FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download 

RUN go build -o bin/main cmd/main.go

ENV PORT=8888

EXPOSE ${PORT}

CMD ["/app/bin/main"]