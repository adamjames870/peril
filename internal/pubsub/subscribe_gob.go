package pubsub

import (
	"bytes"
	"encoding/gob"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeGob[T any](
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

	err := ch.Qos(10, 0, true)
	if err != nil {
		return err
	}

	deliveryChan, errConsume := ch.Consume(qu.Name, "", false, false, false, false, nil)
	if errConsume != nil {
		return errConsume
	}

	go func() {
		for delivery := range deliveryChan {

			buf := bytes.NewBuffer(delivery.Body)
			body := new(T)
			decoder := gob.NewDecoder(buf)
			errDecoder := decoder.Decode(&body)

			var err error
			if errDecoder != nil {
				err = delivery.Nack(false, false)
				fmt.Println("NackDiscard for bad gob")
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
