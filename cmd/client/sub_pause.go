package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subPause(state *clientState) error {

	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, state.userName)

	return pubsub.SubscribeJSON(
		state.conn,
		routing.ExchangePerilDirect,
		queueName,
		routing.PauseKey,
		pubsub.Transient,
		handler_pause(state.gameState))

}
