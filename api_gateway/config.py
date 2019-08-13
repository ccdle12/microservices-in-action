"""Configuration file for the Flask based App."""

import os

class Config():
    ENV = 'production'
    DEBUG = False
    TESTING = False

class DevelopmentConfig(Config):
    ENV = 'development'
    DEBUG = True
    TESTING = True
