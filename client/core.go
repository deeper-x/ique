package client

import (
	"log"

	"github.com/deeper-x/ique/configuration"
	"github.com/streadway/amqp"
)

// Pitcher is the Pitch interface
type Pitcher interface {
	Connect() error
	Disconnect() error
	CreateChannel() error
	DestroyChannel() error
	DeclareQueue(string) (amqp.Queue, error)
	Publish(amqp.Queue, string) error
}

// Pitch todo doc
type Pitch struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// Connect open new Connection
func (p *Pitch) Connect() error {
	conn, err := amqp.Dial(configuration.ConnString)
	p.Connection = conn

	return err
}

// Disconnect close connection
func (p *Pitch) Disconnect() error {
	return p.Connection.Close()
}

// CreateChannel create channel from connection
func (p *Pitch) CreateChannel() error {
	ch, err := p.Connection.Channel()

	if err != nil {
		log.Println("debug", err)
	}

	p.Channel = ch
	return err
}

// DestroyChannel close open channel
func (p *Pitch) DestroyChannel() error {
	return p.Channel.Close()
}

// DeclareQueue sends the message to the queue
func (p *Pitch) DeclareQueue(body string) (amqp.Queue, error) {
	queue, err := p.Channel.QueueDeclare(
		body,
		false,
		false,
		false,
		false,
		nil,
	)

	return queue, err
}

// Publish todo doc
func (p *Pitch) Publish(qu amqp.Queue, msg string) error {
	err := p.Channel.Publish(
		"",
		qu.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	if err != nil {
		return err
	}

	log.Printf("Message sent on %s: %s", qu.Name, msg)

	return nil
}

// Run start pitcher instance
func Run(p Pitcher, name string, msg string) error {
	err := p.Connect()
	if err != nil {
		return err
	}

	err = p.CreateChannel()
	if err != nil {
		return err
	}

	qu, err := p.DeclareQueue(name)
	if err != nil {
		return err
	}

	err = p.Publish(qu, msg)
	if err != nil {
		return err
	}

	err = p.DestroyChannel()
	if err != nil {
		return err
	}

	err = p.Disconnect()
	if err != nil {
		return err
	}

	return nil
}
