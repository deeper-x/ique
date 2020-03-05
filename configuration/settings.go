package configuration

import "fmt"

// QueueName is the declared Name
const QueueName = "mesg-qu"

// RabbitmqURI is the rabbitmq server endpoint
const RabbitmqURI = "localhost:5672"

// ConnString redis connection
var ConnString = fmt.Sprintf("amqp://guest:guest@%s/", RabbitmqURI)

// MonitoredDir directory
var MonitoredDir = "/tmp/monitor"

// AwsTopic is the default topic (is private, will be hidden soon though)
var AwsTopic = "arn:aws:sns:eu-west-3:777350386990:justopic"
