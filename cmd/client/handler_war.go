package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
)

func handlerWar(state *clientState) func(attacker gamelogic.Player) pubsub.AckType {

	return func(attacker gamelogic.Player) pubsub.AckType {
		defer fmt.Print("> ")
		recog := gamelogic.RecognitionOfWar{
			Attacker: attacker,
			Defender: state.gameState.Player,
		}
		warOutcome, _, _ := state.gameState.HandleWar(recog)
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
