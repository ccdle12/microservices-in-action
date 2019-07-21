"""Entry point for the Flask App."""


from flask import Flask
import os

app = Flask(__name__)
app.config.from_object('config')

from app import routes
from app import order_service_pb2
from app import order_service_pb2_grpc
