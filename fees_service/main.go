package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	// Connect the event queue.
	conn, err := amqp.Dial("amqp://admin:Admin123@order_event_queue:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to the event queue.")
	}
	defer conn.Close()

	// Opens a channel to the event queue.
	event_queue_channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel.")
	}

	// Creates a channel to receive closing notifications from the event queue.
	// This means the notify channel will receive a message when the connection
	// to the event queue has been lost.
	notify := conn.NotifyClose(make(chan *amqp.Error)) // error channel

	// Creates a queue `order_placed`, the fees_service will receive
	// messages from the event_queue when an order is placed on the market.
	order_placed_queue, err := event_queue_channel.QueueDeclare(
		// "fees_service_queue", // Queue name.
		"order_placed", // Queue name.
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue.")
	}

	// Consumes and listens to the queue.
	// `order_placed_msgs` is a *amqp.Delivery channel.
	order_placed_msgs, err := event_queue_channel.Consume(
		order_placed_queue.Name, // queue
		"",                      // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     //args
	)

	logger.Info(" [*] Waiting for messages. To exit press CTRL+C")
	// Blocking loop, waits for messages from the channels.
	for {
		select {
		case <-notify:
			log.Printf(" [*] Reconnecting to event queue")
			// NOTE: THIS IS A HACK
			// When receiveing an unresponsive message from the event queue.
			// This will panic and cause the container to restart. This is a
			// bit dirty since we shoudn't rely on the container restarting to
			// enable reconnection.
			panic("Notified that the event queue is unreponsive.")
		case msg := <-order_placed_msgs:
			logger.Info("Received order placed messages: %s", msg)
		}
	}
}
