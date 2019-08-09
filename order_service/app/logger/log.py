"""Log formats the logging output that will be picked up by the elk stack."""


import logging
import sys


logger = logging.getLogger()
logger.setLevel(logging.DEBUG)

handler = logging.StreamHandler(sys.stdout)
handler.setLevel(logging.DEBUG)

formatter = logging.Formatter(
    '%(asctime)s - {} - %(levelname)s - %(message)s'.format('order_service')
)
handler.setFormatter(formatter)

logger.addHandler(handler)
