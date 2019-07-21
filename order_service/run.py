"""
The entrypoint for the order_service program.

Creates the gRPC Server to listen for internal communication
from the api_gateway.
"""

from concurrent.futures import ThreadPoolExecutor
import time
import os
import grpc
from app import order_service_pb2_grpc
from app.order_server import OrderServer


# Create the server.
SERVER = grpc.server(ThreadPoolExecutor(max_workers=10))
order_service_pb2_grpc.add_OrderServicer_to_server(
    OrderServer(),
    SERVER
)

# Open the server on it's port and run the server.
SERVER.add_insecure_port('[::]:%s' % os.environ['ORDER_SERVICE_PORT'])
SERVER.start()

# Keep the program running.
while True:
    time.sleep(10)
