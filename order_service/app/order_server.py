"""
Implementation of the gRPC Server, allows the api_gatway to place orders on
the event_queue.
"""

from app import order_service_pb2_grpc, order_service_pb2, models, db, event_queue_client, utils
from flask_sqlalchemy import SQLAlchemy
import uuid


class OrderServer(order_service_pb2_grpc.OrderServicer):
    """
    OrderServer exposes end points for the api_gateway to place orders on
    the market.
    """
    def GetAllOrders(self, request, context):
        orders = models.Order.query.all()

        def build_order_status(order):
            return order_service_pb2.OrderStatusResponse(
                order_id=order.order_id,
                user_id=order.user_id,
                symbol=order.symbol.upper(),
                amount=order.amount,
                status=utils.tx_status(order.status)
            )

        return order_service_pb2.OrderStatusAllResponse(
            orders=[build_order_status(order) for order in orders]
        )

    def CreateOrder(self, request, context):
        order_id=str(uuid.uuid4())

        # Create the order model.
        new_order = models.Order(
            order_id=order_id,
            user_id="1",
            symbol=request.symbol.upper(),
            amount=request.amount,
            status=0
        )
        db.session.add(new_order)
        db.session.commit()

        # Send Request to the Event Queue.
        try:
            place_order = event_queue_client.EventQueueClient()
            response = place_order.call(str(request))
        except:
            return order_service_pb2.OrderResponse(
                status=utils.tx_status(2)
            )

        # Update the status of the order in the DB.
        order = models.Order.query.filter_by(order_id=order_id).first()
        order.status = int(response)
        db.session.commit()

        return order_service_pb2.OrderResponse(
            status=utils.tx_status(int(response))
        )
