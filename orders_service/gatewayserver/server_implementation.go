package gatewayserver

import (
	proto "github.com/simplebank/orders_service/grpc"
	"golang.org/x/net/context"
)

// Server is the implementation struct for the proto file describing the endpoints
// callable from the api_gateway service.
type Server struct{}

// CreateOrder handles a request from the api_gateway to place an order on the
// market.
func (s *Server) CreateOrder(context.Context, *proto.OrderRequest) (*proto.OrderResponse, error) {
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
