package eventqueue

import (
	"github.com/streadway/amqp"
)

// OrderEventQueue is a struct that wraps a connection to the rabbitmq order
// event queue. It implements QueueClient.
type OrderEventQueue struct {
	Conn       *amqp.Connection
	Chann      *amqp.Channel
	Queue      amqp.Queue
	CloseChann chan *amqp.Error
}

// New is the constructor for an OrderEventQueue.
func New() (*OrderEventQueue, error) {
	conn, err := amqp.Dial("amqp://admin:Admin123@order_event_queue:5672/")
	if err != nil {
		return nil, err
	}

	chann, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Creates a queue if it doesn't already exists, otherwise connects to the
	// existing queue.
	queue, err := chann.QueueDeclare(
		"order_placed", // Queue name.
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	// Channel that receives a close notification from the order event queue.
	closeChann := conn.NotifyClose(
		make(chan *amqp.Error),
	)

	return &OrderEventQueue{
		Conn:       conn,
		Chann:      chann,
		Queue:      queue,
		CloseChann: closeChann}, nil
}

// Cleanup should be used after calling New() as a deferred function. It closes
// any existing connections.
func (o *OrderEventQueue) Cleanup() {
	o.Conn.Close()
}

// SendMessage is the QueueClient implementation, sending a message to the
// created Queue.
func (o *OrderEventQueue) SendMessage(msg []byte) error {
	err := o.Chann.Publish("", o.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         msg,
	})
	if err != nil {
		return err
	}

	return nil
}

// QueueClient implementation for processing a received message.
func (o *OrderEventQueue) ProcessMessage() error {
	return nil
}
