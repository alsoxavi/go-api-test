FROM golang:1.15-alpine3.13

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git

RUN go mod download
RUN go build -o main .

CMD ["/app/main"]