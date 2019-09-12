package main

import (
	"github.com/simplebank/orders_service/gatewayserver"
	"log"
	"time"
)

func main() {
	// Creates a grpc server for the api_gateway to send requests.
	srv, err := gatewayserver.MakeAPIGatewayServer()
	if err != nil {
		log.Fatalf("Failed to create server for the APIGateway: %v", err)
	}
	defer srv.Serve()

	// TEMP
	// Will be used for consuming channels.
	go func() {
		for {
			select {
			default:
				log.Println("blocking...")
				time.Sleep(5 * time.Second)
			}
		}
	}()
}
