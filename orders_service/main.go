package main

import (
	"github.com/simplebank/orders_service/eventqueue"
	"github.com/simplebank/orders_service/gatewaygrpc"
	"log"
	"time"
)

func main() {
	// Creates a connection to the order event queue.
	orderEventQueue, err := eventqueue.New()
	if err != nil {
		log.Fatalf("Failed to create queue client")
	}
	defer orderEventQueue.Cleanup()

	// Creates a grpc server, allowing the api_gateway to send requests to the
	// order service.
	apiGatewayServer, err := gatewaygrpc.New(orderEventQueue)
	if err != nil {
		log.Fatalf("Failed to create server for the APIGateway: %v", err)
	}
	defer apiGatewayServer.Serve() // is blocking.

	// Listen to messages from the Event Queue.
	go func() {
		for {
			select {
			case <-orderEventQueue.CloseChann:
				log.Println("closing...")
			default:
				log.Println("blocking...")
				time.Sleep(5 * time.Second)
			}
		}
	}()
}
