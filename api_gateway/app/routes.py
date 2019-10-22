"""Routes exposes public facing end points to an end user."""


from app import app
from app.grpc import order_service_pb2, order_service_pb2_grpc
from app.logger.log import logger
from flask import jsonify, request
import uuid
import os
import grpc


@app.route('/order', methods=['POST'])
def order():
    req = request.get_json()
    # TODO (ccdle12) use flask Rest Plus to validate.

    try:
        # Maybe wrap this with a OrderServiceClient?
        channel = grpc.insecure_channel(
            'orders_service:{}'.format(os.environ['ORDERS_SERVICE_PORT'])
        )
    except Exception as e:
        return jsonify('error: {}'.format(str(e)))

    try:
        stub = order_service_pb2_grpc.OrderStub(channel)
    except Exception as e:
        return jsonify('error: {}'.format(str(e)))

    try:
        response = stub.CreateOrder(
            timeout=5,
            request=order_service_pb2.OrderRequest(
                user_id=str(req["user_id"]),
                symbol=req['symbol'].upper(),
                order_size=str(req['order_size']),
                price="123"
            )
        )

        logger.info(f"Received response from grpc orders_service: {response}")
        if 'FAILED' == response.status:
            logger.error('error in order request: {}'.format(str(response.status)))
            return jsonify({'error': str(response.status)}), 500
        else:
            logger.info('response to order request: {}'.format(str(response.status)))
            return jsonify({'message': str(response.status)})

        logger.info('response to order request: {}'.format(str(response.status)))
        return jsonify({'message': 'hello'})

    except Exception as e:
        # logger.error('error in api_gateway, order service is unresponsive')
        logger.error(e)
        return jsonify({'message': 'order_service unavailable'}), 500

    return jsonify({'result': 'placeholder'})

@app.route('/order', methods=['GET'])
def all_orders():
    # Utils?
    print('DEBUG: GET ALL ORDERS')
    # TODO (ccdle12): Marshmallow?
    def deserialize_orders(order):
        order_dict = {}
        order_dict['order_id'] = order.order_id
        order_dict['user_id'] = order.user_id
        order_dict['symbol_id'] = order.symbol
        order_dict['amount'] = order.amount
        order_dict['status'] = order.status

        return order_dict

    try:
        channel = grpc.insecure_channel(
            'orders_service:{}'.format(os.environ['ORDERS_SERVICE_PORT'])
        )
        stub = order_service_pb2_grpc.OrderStub(channel)
        response = stub.GetAllOrders(
            order_service_pb2.OrderStatusAllRequest()
        )

        logger.info('response to get all orders: {}'.format(str(response.orders)))
        return jsonify({'orders': [deserialize_orders(order) for order in response.orders]})
    except Exception as error:
        logger.error('error in get orders request: {}'.format(str(error)))
        return jsonify({'message': 'order_service unavailable'}), 500
