package gatewaygrpc

import (
	// "github.com/simplebank/orders_service/eventqueue"
	proto "github.com/simplebank/orders_service/grpc"
	"golang.org/x/net/context"
	"testing"
)

// Mock for the EventQueue interface when creating a Server.
type EventQueueMock struct{}

func (e *EventQueueMock) SendMessage(msg []byte) error {
	return nil
}

func (e *EventQueueMock) ProcessMessage() error {
	return nil
}

// Tests passing Nil into a Server as the interfacec QueueClient should
// NOT crash the grpc server.
func TestNilQueueClient(t *testing.T) {
	server := Server{nil}

	ctx := context.Background()
	order_req := &proto.OrderRequest{UserId: "123", Symbol: "BTC", Amount: 123}

	_, err := server.CreateOrder(ctx, order_req)
	if err != QueueClientNilError {
		t.Errorf("QueueClientNilError should have been raised.")
	}
}

// Tests that we can use a mock for the EventQueue when creating the server.
func TestQueueClientMock(t *testing.T) {
	eventQueue := &EventQueueMock{}
	server := Server{eventQueue}

	ctx := context.Background()
	order_req := &proto.OrderRequest{UserId: "123", Symbol: "BTC", Amount: 123}

	_, err := server.CreateOrder(ctx, order_req)
	if err != nil {
		t.Errorf("No error should have been raised, the event queue is mocked.")
	}
}
