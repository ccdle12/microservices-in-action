# Orders Service

This service is responsible for receiving trade orders from
the api_gateway service and writing the clients orders to
the `orders_db` and then sending an `order_placed` message
to the `event_queue`.

The main implementation for receiving, processing and emitting
the order can found in `./gatewaygrpc/server_implementation.go`.
