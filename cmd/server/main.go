package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	fmt.Println("Starting Peril server...")

	connectionStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		log.Fatal("cant dial amqp", err)
	}
	defer conn.Close()
	fmt.Println("Connected to amqp server")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("cant create channel", err)
	}

	err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
		IsPaused: true,
	})

	fmt.Println("Sended message")

	if err != nil {
		log.Fatal("cant publish message", err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down...")
}
