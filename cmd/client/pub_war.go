package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func publishWar(ch *amqp.Channel, recogWar gamelogic.RecognitionOfWar) error {

	routingKey := fmt.Sprintf("%s.%s", routing.WarRecognitionsPrefix, recogWar.Attacker.Username)

	return pubsub.PublishJSON(
		ch,
		routing.ExchangePerilTopic,
		routingKey,
		recogWar,
	)

}
