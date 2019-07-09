package main

import (
        "github.com/streadway/amqp"
        "log"
)

func main() {
        // Connect the event queue.
        conn, err := amqp.Dial("amqp://admin:Admin123@order_event_queue:5672/")
        if err != nil {
                log.Fatalf("Failed to connect to the event queue.")
        }
        defer conn.Close()

        // Opens a channel to the event queue.
        ch, err := conn.Channel()
        if err != nil {
                log.Fatalf("Failed to create channel.")
        }

        // Creates a channel to receive notifications from the event queue closing
        // it's connection/going offline.
        notify := conn.NotifyClose(make(chan *amqp.Error)) // error channel

        // Creates a Queue for the fees_service_queue, the fees_service will receive
        // messages from the event_queue when an order is confirmed.
        q, err := ch.QueueDeclare(
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

        // Consumes and listens to the queue. `msgs` is a *amqp.Delivery channel.
        msgs, err := ch.Consume(
                q.Name, // queue
                "",     // consumer
                true,   // auto-ack
                false,  // exclusive
                false,  // no-local
                false,  // no-wait
                nil,    //args
        )

        log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
        // Blocking loop, waits for messages from the channels.
        for {
                select {
                case <-notify:
                        log.Printf(" [*] Reconnecting to event queue")
                        // NOTE: THIS IS A HACK
                        // When an unresponsive message from the event queue. This will panic
                        // and cause the container to restart. This is a bit dirty since we
                        // shoudn't rely on the container restarting to enable reconnection.
                        panic("Notified that the event queue is unreponsive.")
                case msg := <-msgs:
                        log.Printf(" [*] Received message: %s", msg)
                }
        }
}
