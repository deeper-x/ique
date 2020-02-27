package configuration

import "fmt"

// QueueName is the declared Name
const QueueName = "mesg-qu"

// RabbitmqURI is the rabbitmq server endpoint
const RabbitmqURI = "localhost:5672"

// ConnString redis connection
var ConnString = fmt.Sprintf("amqp://guest:guest@%s/", RabbitmqURI)
