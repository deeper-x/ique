package client

import (
	"log"

	"github.com/streadway/amqp"
)

// MockPitch is the Pitch Mock
type MockPitch struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewMockQueue returns a amqp.Queue
func NewMockQueue() amqp.Queue {
	return amqp.Queue{Name: "mock", Messages: 1, Consumers: 1}
}

// Connect mocks connection
func (mp *MockPitch) Connect() error {
	return nil
}

// Disconnect close connection
func (mp *MockPitch) Disconnect() error {
	return nil
}

// CreateChannel create channel from connection
func (mp *MockPitch) CreateChannel() error {
	return nil
}

// DestroyChannel close open channel
func (mp *MockPitch) DestroyChannel() error {
	return nil
}

// DeclareQueue sends the message to the queue
func (mp *MockPitch) DeclareQueue(body string) (amqp.Queue, error) {
	var mockQueue = NewMockQueue()
	return mockQueue, nil
}

// Publish todo doc
func (mp *MockPitch) Publish(qu amqp.Queue, msg string) error {
	log.Printf("Mock Message sent ")
	return nil
}
