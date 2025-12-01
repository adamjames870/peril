package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subWar(state *clientState) error {

	queueName := fmt.Sprintf("%s.%s", routing.WarRecognitionsPrefix, state.gameState.Player.Username)
	routingKey := queueName

	return pubsub.SubscribeJSON(
		state.conn,
		routing.ExchangePerilTopic,
		queueName,
		routingKey,
		pubsub.Durable,
		handlerWar(state),
	)

}
