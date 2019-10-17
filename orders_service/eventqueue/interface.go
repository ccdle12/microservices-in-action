package eventqueue

// QueueClient is the interface for accessing the event queue.
type QueueClient interface {
	SendMessage(msg []byte) error
	ProcessMessage() error
}
