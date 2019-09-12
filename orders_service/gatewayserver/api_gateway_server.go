package gatewayserver

import (
	"fmt"
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

// Constructor for the APIGatewayServer, creates a grpc server using the proto
// file implentation and registers the server to a port and address.
func MakeAPIGatewayServer() (*APIGatewayServer, error) {
	listener, err := createListener(createServerURI())
	if err != nil {
		return nil, err
	}

	return &APIGatewayServer{
		URI:      createServerURI(),
		Listener: listener,
		Server:   createGRPCServer()}, nil
}

func createServerURI() string {
	return fmt.Sprintf("0.0.0.0:%s", os.Getenv("ORDERS_SERVICE_PORT"))
}

func createListener(uri string) (net.Listener, error) {
	listener, err := net.Listen("tcp", uri)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func createGRPCServer() *grpc.Server {
	srv := grpc.NewServer()
	proto.RegisterOrderServer(srv, &Server{})

	return srv
}

// Serve is a method that calls the Server (*grpc.Server) to start listening for
// requests using the Listener (net.Listener) in APIGatewayServer.
func (a *APIGatewayServer) Serve() {
	a.Server.Serve(a.Listener)
}
