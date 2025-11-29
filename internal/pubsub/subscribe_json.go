package pubsub

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
	handler func(T),
) error {

	ch, qu, errBind := DeclareAndBind(conn, exchange, queueName, key, queueType)
	if errBind != nil {
		return errBind
	}

	deliveryChan, errConsume := ch.Consume(qu.Name, "", false, false, false, false, nil)
	if errConsume != nil {
		return errConsume
	}

	go func() {
		for delivery := range deliveryChan {
			body := new(T)
			errUnmarshall := json.Unmarshal(delivery.Body, body)
			if errUnmarshall != nil {
				// How to handle this error?
			} else {
				handler(*body)
			}
			errAck := delivery.Ack(false)
			if errAck != nil {
				// How to handle this error
			}
		}
	}()

	return nil

}
