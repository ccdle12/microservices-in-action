"""Entry point for the Flask App."""


from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import os


app = Flask(__name__)

""" Create an instance of the DB. """
db_connection = 'postgresql://{}:{}@{}:{}/{}'.format(
    os.environ['DB_USER'],
    os.environ['DB_PASSWORD'],
    os.environ['DB_NAME'],
    os.environ['DB_PORT'],
    os.environ['DB_NAME']
)

app.config['SQLALCHEMY_DATABASE_URI'] = db_connection
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
db = SQLAlchemy(app)

from app.models import order_model
from app import event_queue_client
from app import utils
from app import order_server
from app.grpc import order_service_pb2
from app.grpc import order_service_pb2_grpc
