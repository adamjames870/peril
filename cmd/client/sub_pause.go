package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func subPause(state *clientState) error {

	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, state.userName)
	handler := handler_pause(state.gameState)
	errSub := pubsub.SubscribeJSON(state.conn, routing.ExchangePerilDirect, queueName, routing.PauseKey, pubsub.Transient, handler)
	if errSub != nil {
		fmt.Println("Failed to subscribe to direct exchange: " + errSub.Error())
	}

	return nil

}
