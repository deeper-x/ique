package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deeper-x/ique/client"
	"github.com/deeper-x/ique/configuration"
	"github.com/deeper-x/ique/myutils"
	"github.com/deeper-x/ique/network"
	"github.com/deeper-x/ique/server"
)

const name = configuration.QueueName

func main() {
	// 1- Pre-checks: services are up&run
	rabbitService := network.Service{URI: configuration.RabbitmqURI}

	if !rabbitService.IsRunning() {
		log.Fatalf("%s not available", rabbitService.URI)
	}

	// 2 - User input
	fmt.Print("Please insert runner [sender/receiver]:")
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

	case inVal == "sender":
		pitch := client.Pitch{}
		err := client.Run(&pitch, name, "demo text")

		myutils.FailsOnError(err, "Failed running sender...")

	default:
		log.Println("not managed...")
	}
}
