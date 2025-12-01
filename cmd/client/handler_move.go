package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
)

func handlerMove(state *clientState) func(move gamelogic.ArmyMove) pubsub.AckType {

	return func(move gamelogic.ArmyMove) pubsub.AckType {

		defer fmt.Print("> ")

		switch state.gameState.HandleMove(move) {
		case gamelogic.MoveOutComeSafe:
			return pubsub.Ack

		case gamelogic.MoveOutcomeMakeWar:

			if move.Player.Username == state.gameState.Player.Username {
				return pubsub.Ack
			}

			recogWar := gamelogic.RecognitionOfWar{
				Attacker: move.Player,
				Defender: state.gameState.Player,
			}

			errPubWar := publishWar(state.publishCh, recogWar)
			if errPubWar != nil {
				return pubsub.NackRequeue
			} else {
				return pubsub.Ack
			}

		case gamelogic.MoveOutcomeSamePlayer:
			return pubsub.NackDiscard

		default:
			return pubsub.NackDiscard
		}
	}

}
