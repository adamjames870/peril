package main

import (
	"fmt"
	"time"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func handlerWar(state *clientState) func(recWar gamelogic.RecognitionOfWar) pubsub.AckType {

	return func(recWar gamelogic.RecognitionOfWar) pubsub.AckType {

		defer fmt.Print("> ")

		warOutcome, winner, loser := state.gameState.HandleWar(recWar)

		switch warOutcome {
		case gamelogic.WarOutcomeNotInvolved:
			return pubsub.NackRequeue

		case gamelogic.WarOutcomeNoUnits:
			return pubsub.NackDiscard

		case gamelogic.WarOutcomeOpponentWon:
			log := routing.GameLog{
				CurrentTime: time.Now(),
				Message:     fmt.Sprintf("%s won a war against %s", winner, loser),
				Username:    state.gameState.GetUsername(),
			}
			if publishGameLog(state, log) != nil {
				return pubsub.NackRequeue
			}
			return pubsub.Ack

		case gamelogic.WarOutcomeYouWon:
			log := routing.GameLog{
				CurrentTime: time.Now(),
				Message:     fmt.Sprintf("%s won a war against %s", winner, loser),
				Username:    state.gameState.GetUsername(),
			}
			if publishGameLog(state, log) != nil {
				return pubsub.NackRequeue
			}
			return pubsub.Ack

		case gamelogic.WarOutcomeDraw:
			log := routing.GameLog{
				CurrentTime: time.Now(),
				Message:     fmt.Sprintf("A war between %s and %s resulted in a draw", winner, loser),
				Username:    state.gameState.GetUsername(),
			}
			if publishGameLog(state, log) != nil {
				return pubsub.NackRequeue
			}
			return pubsub.Ack

		default:
			fmt.Println("WarOutcome not recognised: ", warOutcome)
			return pubsub.NackDiscard
		}
	}

}
