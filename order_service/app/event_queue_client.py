import pika
import uuid
import os

class EventQueueClient(object):
    # Private Variables.
    # The connection to the order_event_queue.
    _connection = None

    # Channel setup to the order_event_queue.
    _channel = None

    # The queue being used for the placement of orders.
    _callback_queue = None

    # The unique id sent to the rpc_queue, this will identify the source of the
    # request for the server.
    _corr_id = None

    # The response from the server.
    _response = None

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
            self._connection = pika.BlockingConnection(params)
        except pika.exceptions.ConnectionClosed:
            raise Exception()

        self._channel = self._connection.channel()

        # This channel will use the order_placed queue to receive confirmation
        # of the order.
        result = self._channel.queue_declare(queue='order_placed')
        self._callback_queue = result.method.queue

        # Listen on the callback queue and execute `on_response` method when
        # a message is received on the callback queue.
        self._channel.basic_consume(
            queue=self._callback_queue,
            on_message_callback=self.on_response,
            auto_ack=True
        )

    def on_response(self, ch, method, props, body):
        # On response on the callback queue, it checks if the message
        # correlates with the clients request.
        if self._corr_id == props.correlation_id:
            self._response = body
            self._connection.close()

    def call(self, n):
        self._corr_id = str(uuid.uuid4())

        self._channel.basic_publish(
            exchange='',
            routing_key='order_created',
            # correlation id is set to identify the sender.
            # reply_to is set to the anonymous callback queue.
            properties=pika.BasicProperties(
                reply_to=self._callback_queue,
                correlation_id=self._corr_id,
            ),
            body=str(n)
        )

        while self._response is None:
            self._connection.process_data_events()

        return int(self._response)
