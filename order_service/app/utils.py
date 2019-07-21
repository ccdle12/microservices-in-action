"""
Util functions that have singular purpose and/or does not mutate/affect another 
other running service.
"""


STATUS_DICT = {0: 'PENDING', 1: 'CONFIRMED', 2: 'FAILED'}

def tx_status(status_code):
    """Returns the string value of a status code from orders placed."""
    return STATUS_DICT[status_code]
