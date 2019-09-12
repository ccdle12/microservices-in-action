#!/bin/sh

protoc -I ./ order_service.proto --go_out=plugins=grpc:.
