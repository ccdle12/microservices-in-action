"""Routes exposes public facing end points to an end user."""


from app import app
from flask import jsonify, request
import uuid
import os
import grpc
from app.grpc import order_service_pb2, order_service_pb2_grpc

import sys

sys.path.append('..')
from logger.log import logger


@app.route('/order', methods=['POST'])
def order():
    print('DEBUG: POST ORDER')
    req = request.get_json()
    # TODO (ccdle12) use flask Rest Plus to validate.
    if 'amount' not in req:
        return jsonify({'message': 'amount is missing in request'}), 404

    if 'symbol' not in req:
        return jsonify({'message': 'symbol is missing in request'}), 404

    try:
        # Maybe wrap this with a OrderServiceClient?
        channel = grpc.insecure_channel(
            'order_service:{}'.format(os.environ['ORDER_SERVICE_PORT'])
        )
        stub = order_service_pb2_grpc.OrderStub(channel)
        response = stub.CreateOrder(
            order_service_pb2.OrderRequest(
                user_id="123",
                symbol=req['symbol'].upper(),
                amount=req['amount']
            )
        )

        if 'FAILED' == response.status:
            logger.error('error in order request: {}'.format(str(response.status)))
            return jsonify({'error': str(response.status)}), 500
        else:
            logger.info('response to order request: {}'.format(str(response.status)))
            return jsonify({'message': str(response.status)})

        logger.info('response to order request: {}'.format(str(response.status)))
        return jsonify({'message': 'hello'})

    except Exception as error:
        logger.error('error in api_gateway, order service is unresponsive')
        return jsonify({'message': 'order_service unavailable'}), 500

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
            'order_service:{}'.format(os.environ['ORDER_SERVICE_PORT'])
        )
        stub = order_service_pb2_grpc.OrderStub(channel)
        response = stub.GetAllOrders(
            order_service_pb2.OrderStatusAllRequest()
        )

        logger.info('response to get all orders: {}'.format(str(response.orders)))
        return jsonify({'orders': [deserialize_orders(order) for order in response.orders]})
    except:
        logger.error('error in get orders request: {}'.format(str(response)))
        return jsonify({'message': 'order_service unavailable'}), 500
