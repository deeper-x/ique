package server

import (
	"log"

	"github.com/deeper-x/ique/configuration"
	"github.com/streadway/amqp"
)

// Executor is the agent/sender interface
type Executor interface {
	Connect() error
	Disconnect() error
	CreateChannel() error
	DestroyChannel() error
	DeclareQueue(string) (amqp.Queue, error)
	StartServer(amqp.Queue) error
}

// Agent is the server core object
type Agent struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// Connect open new connection to amqp
func (a *Agent) Connect() error {
	conn, err := amqp.Dial(configuration.ConnString)
	a.Connection = conn

	return err
	// myutils.FailsOnError(err, "Failed to connect to rabbitmq...")
}

// Disconnect close amqp connection
func (a *Agent) Disconnect() error {
	return a.Connection.Close()
}

// CreateChannel create channel from connection
func (a *Agent) CreateChannel() error {
	ch, err := a.Connection.Channel()

	a.Channel = ch
	return err
}

// DestroyChannel close open channel
func (a *Agent) DestroyChannel() error {
	return a.Channel.Close()
}

// DeclareQueue sends the message to the queue
func (a *Agent) DeclareQueue(name string) (amqp.Queue, error) {
	queue, err := a.Channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	return queue, err
}

// StartServer server instance
func (a *Agent) StartServer(qu amqp.Queue) error {
	msgs, err := a.Channel.Consume(qu.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	rbChan := make(chan bool)

	go func() {
		for m := range msgs {
			log.Printf("Received: %s", m.Body)
		}
	}()

	log.Println("Waiting for messages....")

	<-rbChan
	return nil
}

// Run start receiver instance
func Run(e Executor, name string) error {
	err := e.Connect()
	if err != nil {
		return err
	}
	err = e.CreateChannel()
	if err != nil {
		return err
	}

	qu, err := e.DeclareQueue(name)
	if err != nil {
		return err
	}

	err = e.StartServer(qu)
	if err != nil {
		return err
	}

	err = e.DestroyChannel()
	if err != nil {
		return err
	}

	err = e.Disconnect()
	if err != nil {
		return err
	}

	return nil
}
