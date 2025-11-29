package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subMoves(state *clientState) error {

	queueName := fmt.Sprintf("%s.%s", routing.ArmyMovesPrefix, state.userName)
	routingKey := fmt.Sprintf("%s.*", routing.ArmyMovesPrefix)

	return pubsub.SubscribeJSON(
		state.conn,
		routing.ExchangePerilTopic,
		queueName,
		routingKey,
		pubsub.Transient,
		handler_move(state.gameState),
	)

}
