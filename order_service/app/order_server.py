"""
Implementation of the gRPC Server, allows the api_gatway to place orders on
the event_queue.
"""


from app import db, utils
from app.grpc import order_service_pb2_grpc, order_service_pb2
from app.models import order_model
from app.event_queue_client import EventQueueClient
from flask_sqlalchemy import SQLAlchemy
import uuid


class OrderServer(order_service_pb2_grpc.OrderServicer):
    """
    OrderServer exposes end points for the api_gateway to place orders on
    the market.
    """
    def GetAllOrders(self, request, context):
        orders = order_model.Order.query.all()

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
        print('DEBUG: CREATE ORDER CALLED')
        order_id=str(uuid.uuid4())

        # Create the order model.
        new_order = order_model.Order(
            order_id=order_id,
            user_id="1",
            symbol=request.symbol.upper(),
            amount=request.amount,
            status=0
        )

        try:
            db.session.add(new_order)
            db.session.commit()
        except:
            return order_service_pb2.OrderResponse(
                status=utils.tx_status(2)
            )

        # Send Request to the Event Queue.
        # TODO (ccdle12): response status should match an enum in protofile.
        try:
            place_order = EventQueueClient()
            response = place_order.call(str(request))
        except:
            return order_service_pb2.OrderResponse(
                status=utils.tx_status(2)
            )

        return order_service_pb2.OrderResponse(
            status=utils.tx_status(1)
        )
