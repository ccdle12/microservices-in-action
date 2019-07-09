from concurrent.futures import ThreadPoolExecutor
import time
import os
import grpc
from app import order_service_pb2_grpc
from app.order_server import OrderServer

# Create the server.
# server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
server = grpc.server(ThreadPoolExecutor(max_workers=10))
order_service_pb2_grpc.add_OrderServicer_to_server(
    OrderServer(),
    server
)

# Open the server on it's port and run the server.
server.add_insecure_port('[::]:%s' % os.environ['ORDER_SERVICE_PORT'])
server.start()

# Keep the program running.
while True:
    time.sleep(10)
