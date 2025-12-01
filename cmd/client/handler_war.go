package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
)

func handlerWar(state *clientState) func(recWar gamelogic.RecognitionOfWar) pubsub.AckType {

	return func(recWar gamelogic.RecognitionOfWar) pubsub.AckType {

		defer fmt.Print("> ")

		warOutcome, _, _ := state.gameState.HandleWar(recWar)

		switch warOutcome {
		case gamelogic.WarOutcomeNotInvolved:
			return pubsub.NackRequeue

		case gamelogic.WarOutcomeNoUnits:
			return pubsub.NackDiscard

		case gamelogic.WarOutcomeOpponentWon:
			return pubsub.Ack

		case gamelogic.WarOutcomeYouWon:
			return pubsub.Ack

		case gamelogic.WarOutcomeDraw:
			return pubsub.Ack

		default:
			fmt.Println("WarOutcome not recognised: ", warOutcome)
			return pubsub.NackDiscard
		}
	}

}
