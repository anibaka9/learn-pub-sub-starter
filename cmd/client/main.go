package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	connectionStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		log.Fatal("cant dial amqp", err)
	}
	defer conn.Close()
	fmt.Println("Connected to amqp server")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatal("cant get username", err)
	}
	
}
