"""
Market Service listens to order_created events, and then
places the orders on the market.
"""
import pika
import time
import os

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
connection = pika.BlockingConnection(params)
channel = connection.channel()

# An rpc_queue is created by the server. This will be used by the client
# to send requests.
channel.queue_declare(queue='order_created')

def send_order(order):
    # Sleep to simulate another RPC call to an exchange.
    time.sleep(2)
    return 1

def on_request(ch, method, props, body):
    """
      on_request is the function called when receiving a request from the client
    """
    # The body is assumed to be able to be cast to an int.
    print(" [.] order(%s)" % str(body))

    # Call the send_order function to place order on the market.
    response = send_order(str(body))

    # Send a message to the Fee Service Queue.
    publish_to_fee_service(str(body))

    # correlation_id this match the response with a request.
    properties = pika.BasicProperties(correlation_id=props.correlation_id)
    ch.basic_publish(exchange='',
                     routing_key=props.reply_to,
                     properties=properties,
                     body=str(response)
                     )
    ch.basic_ack(delivery_tag=method.delivery_tag)

def publish_to_fee_service(msg):
    """ Publishes a message to the fee_service_queue """
    channel.basic_publish(
        exchange='',
        routing_key='order_placed',
        body=msg
    )

# Listen on the `order_created` for messages for the market service.
channel.basic_consume(queue='order_created', on_message_callback=on_request)

print(" [x] Awaiting RPC requests")
channel.start_consuming()
