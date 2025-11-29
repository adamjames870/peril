package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func handlerPause(gs *gamelogic.GameState) func(state routing.PlayingState) pubsub.AckType {

	return func(state routing.PlayingState) pubsub.AckType {
		defer fmt.Print("> ")
		gs.HandlePause(state)
		return pubsub.Ack
	}

}
