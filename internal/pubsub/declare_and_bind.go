package pubsub

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	Durable SimpleQueueType = iota
	Transient
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {

	ch, errChannel := conn.Channel()
	if errChannel != nil {
		return nil, amqp.Queue{}, errors.New("Failed to open channel: " + errChannel.Error())
	}

	durable := queueType == Durable
	autoDelete := queueType == Transient
	exclusive := autoDelete

	queue, errQueue := ch.QueueDeclare(queueName, durable, autoDelete, exclusive, false, nil)
	if errQueue != nil {
		return nil, amqp.Queue{}, errQueue
	}

	qrrBind := ch.QueueBind(queue.Name, key, exchange, false, nil)
	if qrrBind != nil {
		return nil, amqp.Queue{}, qrrBind
	}

	return ch, queue, nil

}
