package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType string

const (
	Durable   SimpleQueueType = "durable"
	Transient SimpleQueueType = "transient"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	msgBody, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("cant marshal val: %w", err)
	}
	err = ch.PublishWithContext(context.Background(), exchange, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
	})
	if err != nil {
		return fmt.Errorf("cant publish message: %w", err)
	}
	return nil
}

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("cant create channel %w", err)
	}

	q, err := ch.QueueDeclare(queueName, queueType == Durable, queueType == Transient, queueType == Transient, false, nil)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, fmt.Errorf("cant create queue %w", err)
	}
	err = ch.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, fmt.Errorf("cant bind queue to exchang %w", err)
	}
	return ch, q, nil
}
