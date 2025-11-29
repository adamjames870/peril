package pubsub

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
	handler func(T) AckType,
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
			var err error
			if errUnmarshall != nil {
				err = delivery.Nack(false, false)
				fmt.Println("NackDiscard for bad json")
			} else {
				switch handler(*body) {
				case Ack:
					err = delivery.Ack(false)
					fmt.Println("Ack")
				case NackRequeue:
					err = delivery.Nack(false, true)
					fmt.Println("NackRequeue")
				case NackDiscard:
					err = delivery.Nack(false, false)
					fmt.Println("NackDiscard")
				}
			}
			if err != nil {
				fmt.Println("Error in ack/nack: " + err.Error())
			}
		}
	}()

	return nil

}
