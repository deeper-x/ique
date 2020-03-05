// # iQue - Description

// Listening for file creation in a monitored directory, producing:
// - content parsing
// - reading and sending to local receiver
// - notifing SQS w/ SNS

// Basically is a typical Producer->Consumer model with queue, handled by a monitoring agent. AWS integration allows to publish to a queue SQS with subscription on SNS topic.

// ## AWS SQS<SNS integration details

// In order to build a public service, resource content is delivered to distributed applications with polling model SQS/SNS on AWS as well: given a notifier's topic [AWS::SNS], registered queue service [AWS::SQS] subscription allows message to be avaible for polling, decoupling sending & receiving, not requiring to be concurrently available and persisting for later time consumption.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deeper-x/ique/configuration"
	"github.com/deeper-x/ique/filesys"
	"github.com/deeper-x/ique/myutils"
	"github.com/deeper-x/ique/network"
	"github.com/deeper-x/ique/server"
)

const name = configuration.QueueName

func main() {
	// 1- Pre-checks: services are up&run
	rabbitService := network.Service{URI: configuration.RabbitmqURI}

	if !rabbitService.IsRunning() {
		log.Fatalf("Rabbitmq unavailable on %s", rabbitService.URI)
	}

	// 2 - User input
	fmt.Print("Please insert runner [receiver/listener]:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	inVal := strings.Replace(input, "\n", "", -1)

	// 3 - Action routing
	switch {
	case inVal == "receiver":
		agent := server.Agent{}
		err = server.Run(&agent, name)

		myutils.FailsOnError(err, "Failed running receiver...")

	case inVal == "listener":
		fileManager := filesys.FileManager{Pwd: configuration.MonitoredDir}
		err = filesys.RunListen(fileManager)

		myutils.FailsOnError(err, "Failed running listener...")

	default:
		log.Println("not managed...")
	}

}
