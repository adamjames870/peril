package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subWar(state *clientState) error {

	queueName := routing.WarRecognitionsPrefix
	routingKey := fmt.Sprintf("%s.*", routing.WarRecognitionsPrefix)

	return pubsub.SubscribeJSON(
		state.conn,
		routing.ExchangePerilTopic,
		queueName,
		routingKey,
		pubsub.Durable,
		handlerWar(state),
	)

}
