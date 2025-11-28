package pubsub

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {

	jsn, errJsn := json.Marshal(val)
	if errJsn != nil {
		return errJsn
	}

	payload := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsn,
	}

	err := ch.PublishWithContext(context.Background(), exchange, key, false, false, payload)
	if err != nil {
		return err
	}

	return nil

}
