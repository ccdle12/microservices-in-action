FROM golang:alpine

RUN apk add git \
    && go get github.com/streadway/amqp \
    && go get github.com/pilu/fresh \
    && go get github.com/sirupsen/logrus

WORKDIR /go/src/github.com/simplebank/fees_service
COPY ./ ./

# Only for dev environment.
CMD ["fresh"]
