from app import db

class Order(db.Model):
    """ Model representing an Order Places by a user.
      Params:
        order_id: A unique id generated the place order.
        user_id: The unique id of the user placing the order (TODO: Link it to the user table)
        symbol: A symbol for the stock.
        amount: The number of orders placed.
        status: A integer representing the state.
          0 = Pending
          1 = Confirmed
          2 = Failed
    """
    order_id = db.Column(db.String(50), unique=True, primary_key=True)
    user_id = db.Column(db.String(50))
    symbol = db.Column(db.String(50))
    amount = db.Column(db.Integer)
    status = db.Column(db.Integer)
