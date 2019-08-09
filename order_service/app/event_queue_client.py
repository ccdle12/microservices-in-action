"""
Event Queue Client allows the order_service to emit events
to the order_event_queue.
"""


import pika
import uuid
import os


class EventQueueClient(object):
    # The connection to the order_event_queue.
    connection = None

    # Channel to the order_event_queue.
    channel = None

    # The queue being used for the placement of orders.
    callback_queue = None

    # The unique id sent to the rpc_queue, this will identify the source of the
    # request for the server.
    corr_id = None

    # The response from the server.
    response = None

    def __init__(self):
        credentials = pika.PlainCredentials(
            os.environ['ORDER_EQ_USER'],
            os.environ['ORDER_EQ_PASSWORD']
        )

        params = pika.ConnectionParameters(
            host='order_event_queue',
            port=os.environ['ORDER_EQ_PORT'],
            virtual_host='/',
            credentials=credentials
        )

        try:
            self.connection = pika.BlockingConnection(params)
            self.channel = self.connection.channel()
        except pika.exceptions.ConnectionClosed:
            raise Exception()

        # This channel will use the order_placed queue to receive confirmation
        # of the order.
        result = self.channel.queue_declare(queue='order_placed')
        self.callback_queue = result.method.queue

        # Listen on the callback queue and execute `on_response` method when
        # a message is received on the callback queue.
        self.channel.basic_consume(
            queue=self.callback_queue,
            on_message_callback=self.on_response,
            auto_ack=True
        )

    def on_response(self, ch, method, props, body):
        # On response on the callback queue, it checks if the message
        # correlates with the clients request.
        if self.corr_id == props.correlation_id:
            self.response = body
            self.connection.close()

    def call(self, n):
        self.corr_id = str(uuid.uuid4())

        self.channel.basic_publish(
            exchange='',
            routing_key='order_created',
            # correlation id is set to identify the sender.
            # reply_to is set to the anonymous callback queue.
            properties=pika.BasicProperties(
                reply_to=self.callback_queue,
                correlation_id=self.corr_id,
            ),
            body=str(n)
        )

        while self.response is None:
            self.connection.process_data_events()

        return int(self.response)
