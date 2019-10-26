package gatewaygrpc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/simplebank/orders_service/eventqueue"
	proto "github.com/simplebank/orders_service/grpc"
	"github.com/simplebank/orders_service/models"
	"golang.org/x/net/context"
)

var (
	QueueClientNilError = errors.New("QueueClient is nil")
	QueueSendMsgError   = errors.New("Failed to send message to queue")
)

// Server is the implementation struct for the proto file describing the endpoints
// callable from the api_gateway service.
type Server struct {
	QueueClient eventqueue.QueueClient
}

// CreateOrder handles a request from the api_gateway to place an order on the
// market.
func (s *Server) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	if s.QueueClient == nil {
		return nil, QueueClientNilError
	}

	// 1. Write the order to the DB.
	//  - Fail return the user it has failed.

	// 2. Send a message to the order event queue as a order has been received.
	err := s.QueueClient.SendMessage([]byte("Hello"))
	if err != nil {
		// TODO(ccdle12): Log the internal technical error.
		return nil, err
	}

	// 3. Respond to the the user that everything was ok.
	return &proto.OrderResponse{Status: "Order Placed"}, nil
}

// GetAllOrders returns a list of all orders for a particular user.
func (s *Server) GetAllOrders(context.Context, *proto.OrderStatusAllRequest) (*proto.OrderStatusAllResponse, error) {
	orderStatusResponse := &proto.OrderStatusResponse{
		OrderId: "1",
		UserId:  "1",
		Symbol:  "BTC",
		Amount:  134,
		Status:  "Pending",
	}
	allResponses := []*proto.OrderStatusResponse{orderStatusResponse}

	return &proto.OrderStatusAllResponse{Orders: allResponses}, nil
}

// OrderReqToClientOrder is a hacky... temporary solution to convert a
// proto.OrderRequest to a models.ClientOrder. The reason for the function
// being here is that neither object should really know about each other in a
// concrete sense. I would prefer a more elegant solution in the future.
// Equivalent to a Rust::From trait.
func OrderReqToClientOrder(order_req *proto.OrderRequest) (*models.ClientOrder, error) {
	orderId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &models.ClientOrder{
		Id:        orderId,
		Symbol:    order_req.Symbol,
		OrderSize: order_req.OrderSize,
		Price:     order_req.Price,
	}, nil
}
