package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

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
	_, _, err = pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, fmt.Sprintf("%s.%s", routing.PauseKey, username), routing.PauseKey, pubsub.Transient)
	if err != nil {
		log.Fatal("cant declare and bind channel to queue", err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down...")
}
