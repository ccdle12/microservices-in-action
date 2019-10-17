package gatewaygrpc

import (
	"fmt"
	"github.com/simplebank/orders_service/eventqueue"
	proto "github.com/simplebank/orders_service/grpc"
	"google.golang.org/grpc"
	"net"
	"os"
)

// APIGatewayServer is the main server and entry point for creating a server that
// allows the api_gateway to send requests.
type APIGatewayServer struct {
	URI      string
	Listener net.Listener
	Server   *grpc.Server
}

// New constructs the APIGatewayServer, creates a grpc server using the proto
// file implentation and registers the server to a port and address.
func New(queue eventqueue.QueueClient) (*APIGatewayServer, error) {
	listener, err := net.Listen("tcp", serverURI())
	if err != nil {
		return nil, err
	}

	return &APIGatewayServer{
		URI:      serverURI(),
		Listener: listener,
		Server:   grpcServer(queue)}, nil
}

// TODO(ccdle12): port server could be none.
func serverURI() string {
	return fmt.Sprintf("0.0.0.0:%s", os.Getenv("ORDERS_SERVICE_PORT"))
}

func grpcServer(queue eventqueue.QueueClient) *grpc.Server {
	server := grpc.NewServer()
	proto.RegisterOrderServer(server, &Server{queue})

	return server
}

// Serve is a method that calls the Server (*grpc.Server) to start listening for
// tcp requests using the Listener (net.Listener).
func (a *APIGatewayServer) Serve() {
	a.Server.Serve(a.Listener)
}
