package pubsub

import (
	"bytes"
	"context"
	"encoding/gob"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishGob[T any](ch *amqp.Channel, exchange, key string, val T) error {

	var gb bytes.Buffer
	encoder := gob.NewEncoder(&gb)

	errGob := encoder.Encode(val)
	if errGob != nil {
		return errGob
	}

	payload := amqp.Publishing{
		ContentType: "application/gob",
		Body:        gb.Bytes(),
	}

	err := ch.PublishWithContext(context.Background(), exchange, key, false, false, payload)
	if err != nil {
		return err
	}

	return nil

}
