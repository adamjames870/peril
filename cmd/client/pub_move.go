package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func publishMove(ch *amqp.Channel, userName string, move *gamelogic.ArmyMove) error {

	routingKey := fmt.Sprintf("%s.%s", routing.ArmyMovesPrefix, userName)

	err := pubsub.PublishJSON(
		ch,
		routing.ExchangePerilTopic,
		routingKey,
		move,
	)

	if err != nil {
		return err
	}

	return nil

}
