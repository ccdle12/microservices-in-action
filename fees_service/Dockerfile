FROM golang:alpine

RUN apk add git \
    && go get github.com/streadway/amqp

WORKDIR /go/src/github.com/simplebank/fees_service
COPY ./ ./

CMD ["go", "run", "main.go"]
