package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func handlerGameLogs(state *serverState) func(log routing.GameLog) pubsub.AckType {

	return func(log routing.GameLog) pubsub.AckType {
		defer fmt.Print("> ")
		err := gamelogic.WriteLog(log)
		if err != nil {
			return pubsub.NackDiscard
		}
		return pubsub.Ack
	}

}
