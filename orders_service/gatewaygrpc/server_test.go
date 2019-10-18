package gatewaygrpc

import (
	"errors"
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

// Failed Mock for the EventQueue interface when creating a Server.
type FailedEventQueueMock struct{}

func (f *FailedEventQueueMock) SendMessage(msg []byte) error {
	return QueueSendMsgError
}

func (f *FailedEventQueueMock) ProcessMessage() error {
	return errors.New("Mock failed process message")
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

// Test that the raised error should have been a QueueSendMsgError.
func TestQueueSendMsgError(t *testing.T) {
	eventQueue := &FailedEventQueueMock{}
	server := Server{eventQueue}

	ctx := context.Background()
	order_req := &proto.OrderRequest{UserId: "123", Symbol: "BTC", Amount: 123}

	_, err := server.CreateOrder(ctx, order_req)
	if err != QueueSendMsgError {
		t.Errorf("Error raised should have been a QueueSendMsgError: receive: %v", err)
	}
}
