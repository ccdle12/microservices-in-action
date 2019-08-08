#!/bin/sh

# python -m grpc_tools.protoc -I./ --python_out=../app --grpc_python_out=../app ./order_service.proto
python -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. ./order_service.proto
