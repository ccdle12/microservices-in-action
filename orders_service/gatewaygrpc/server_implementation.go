package gatewaygrpc

import (
	"errors"
	"github.com/simplebank/orders_service/eventqueue"
	proto "github.com/simplebank/orders_service/grpc"
	"golang.org/x/net/context"
)

var (
	QueueClientNilError = errors.New("QueueClient is nil.")
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
	s.QueueClient.SendMessage([]byte("Hello"))

	// 2. Send a message to the order event queue as a order has been received.
	// 3. Respond the the user that everything was ok.
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
