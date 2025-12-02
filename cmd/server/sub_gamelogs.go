package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subGameLogs(state *serverState) error {
	queueName := routing.GameLogSlug
	routingKey := fmt.Sprintf("%s.*", routing.GameLogSlug)

	return pubsub.SubscribeGob(
		state.conn,
		routing.ExchangePerilTopic,
		queueName,
		routingKey,
		pubsub.Durable,
		handlerGameLogs(state),
	)
}
