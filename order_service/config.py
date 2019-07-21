"""Configuration file for the Flask based App."""
import os

class Config():
    ENV = 'production'
    DEBUG = False
    TESTING = False
    PORT = os.environ['ORDER_SERVICE_PORT']

class DevelopmentConfig(Config):
    ENV = 'development'
    DEBUG = True
    TESTING = True
