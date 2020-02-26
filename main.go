package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deeper-x/ique/client"
	"github.com/deeper-x/ique/configuration"
	"github.com/deeper-x/ique/server"
)

const name = configuration.QueueName

func main() {
	fmt.Print("Please insert runner [sender/receiver]:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	inVal := strings.Replace(input, "\n", "", -1)

	switch {
	case inVal == "receiver":
		agent := server.Agent{}
		server.Run(&agent, name)

	case inVal == "sender":
		pitch := client.Pitch{}
		client.Run(&pitch, name, "demo text")

	default:
		log.Println("not managed...")
	}
}
