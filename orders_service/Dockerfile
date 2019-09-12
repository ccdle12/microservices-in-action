FROM golang:alpine

RUN apk add git \
    && go get github.com/streadway/amqp \
    && go get github.com/pilu/fresh \
    && go get github.com/sirupsen/logrus \
    && go get google.golang.org/grpc

WORKDIR /go/src/github.com/simplebank/orders_service
COPY ./ ./

# Only for dev environment.
CMD ["fresh"]
