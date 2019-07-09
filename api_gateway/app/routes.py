from app import app
from flask import jsonify, request
import uuid
import os

import grpc
from app import order_service_pb2
from app import order_service_pb2_grpc

@app.route('/order', methods=['POST'])
def order():
    data = request.get_json()

    if 'amount' not in data:
        return jsonify({'message': 'amount is missing in request'}), 404

    if 'symbol' not in data:
        return jsonify({'message': 'symbol is missing in request'}), 404

    try:
        channel = grpc.insecure_channel('order_service:%s' % os.environ['ORDER_SERVICE_PORT'])
        stub = order_service_pb2_grpc.OrderStub(channel)
        response = stub.CreateOrder(
            order_service_pb2.OrderRequest(
                user_id="123",
                symbol=data['symbol'].upper(),
                amount=data['amount'])
        )

        return jsonify({'message': str(response.status)})
    except:
        return jsonify({'message': 'order_service unavailable'}), 500

@app.route('/order', methods=['GET'])
def all_orders():
    # Utils?
    def deserialize_orders(order):
        order_dict = {}
        order_dict['order_id'] = order.order_id
        order_dict['user_id'] = order.user_id
        order_dict['symbol_id'] = order.symbol
        order_dict['amount'] = order.amount
        order_dict['status'] = order.status

        return order_dict

    try:
        channel = grpc.insecure_channel('order_service:%s' % os.environ['ORDER_SERVICE_PORT'])
        stub = order_service_pb2_grpc.OrderStub(channel)
        response = stub.GetAllOrders(
            order_service_pb2.OrderStatusAllRequest()
        )

        return jsonify({'orders': [deserialize_orders(order) for order in response.orders]})
    except:
        return jsonify({'message': 'order_service unavailable'}), 500
