package server

import (
	"log"

	"github.com/streadway/amqp"
)

// MockAgent is the server core object
type MockAgent struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewMockQueue returns a amqp.Queue
func NewMockQueue() amqp.Queue {
	return amqp.Queue{Name: "mock", Messages: 1, Consumers: 1}
}

// Connect open new connection to amqp
func (a *MockAgent) Connect() error {
	return nil
}

// Disconnect close amqp connection
func (a *MockAgent) Disconnect() error {
	return nil
}

// CreateChannel create channel from connection
func (a *MockAgent) CreateChannel() error {
	return nil
}

// DestroyChannel close open channel
func (a *MockAgent) DestroyChannel() error {
	return nil
}

// DeclareQueue sends the message to the queue
func (a *MockAgent) DeclareQueue(name string) (amqp.Queue, error) {
	queue := NewMockQueue()

	return queue, nil
}

// StartServer server instance
func (a *MockAgent) StartServer(qu amqp.Queue) error {
	log.Println("Waiting for Mock messages")
	return nil
}
