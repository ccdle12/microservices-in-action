FROM golang:alpine

RUN apk add git \
    && go get github.com/streadway/amqp \
    && go get github.com/pilu/fresh \
    && go get github.com/sirupsen/logrus \
    && go get github.com/jinzhu/gorm \
    && go get google.golang.org/grpc \
    && go get github.com/lib/pq \
	&& go get github.com/google/uuid

WORKDIR /go/src/github.com/simplebank/orders_service
COPY ./ ./

# Only for dev environment.
CMD ["fresh"]
