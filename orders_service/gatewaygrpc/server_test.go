package gatewaygrpc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
	order_req := &proto.OrderRequest{
		UserId:    "123",
		Symbol:    "BTC",
		OrderSize: "123",
		Price:     "123",
	}

	_, err := server.CreateOrder(ctx, order_req)
	if err != QueueClientNilError {
		t.Errorf("QueueClientNilError should have been raised.")
	}
}

// Tests that we can convert a proto object to json.
func TestOrderRequsetToClientOrder(t *testing.T) {
	order_req := &proto.OrderRequest{
		UserId:    "123",
		Symbol:    "BTC",
		OrderSize: "123",
		Price:     "123",
	}
	client_order := OrderReqToClientOrder(order_req)

	// Test that all of the field values in client_order match the order_req.
	var testCases = []struct {
		input    string
		expected string
	}{
		{client_order.Id, order_req.UserId},
		{client_order.Symbol, order_req.Symbol},
		{client_order.OrderSize, order_req.OrderSize},
		{client_order.Price, order_req.Price},
	}

	for _, test := range testCases {
		if test.input != test.expected {
			t.Errorf(
				"test failed: input: %v, expected: %v", test.input, test.expected,
			)
		}
	}
}

// Tests that the server implementation can write a ClientOrder to the DB.
func TestWriteClientOrder(t *testing.T) {
	orderReq := &proto.OrderRequest{
		UserId:    "123",
		Symbol:    "BTC",
		OrderSize: "123",
		Price:     "123",
	}
	clientOrder := OrderReqToClientOrder(orderReq)

	// TODO(ccdle12): move this to a dbclient?
	dbURI := fmt.Sprintf("host=order_db user=order_db dbname=order_db sslmode=disable password=some_password")
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		t.Errorf("Creating the connection to the DB failed: %v", err)
	}
	defer db.Close()

	// db.NewRecord(clientOrder)
	db.AutoMigrate(&clientOrder)
	db.Create(&clientOrder)
	// db.Save(&clientOrder)
}
