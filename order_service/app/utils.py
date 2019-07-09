STATUS_DICT = {0: 'PENDING', 1: 'CONFIRMED', 2: 'FAILED'}

def tx_status(status_code):
    return STATUS_DICT[status_code]
