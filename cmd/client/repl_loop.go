package main

import (
	"errors"
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
)

func ReplLoop(state *clientState) {

	for {
		words := gamelogic.GetInput()
		shouldQuit := false

		if len(words) == 0 {
			continue
		}

		var err error
		switch GetClientCommand(words[0]) {
		case Spawn:
			err = cmdSpawn(state, words)

		case Move:
			err = cmdMove(state, words)

		case Status:
			state.gameState.CommandStatus()

		case Help:
			gamelogic.PrintClientHelp()

		case Spam:
			err = cmdSpam(state, words)

		case Quit:
			gamelogic.PrintQuit()
			shouldQuit = true

		default:
			err = errors.New("Unknown command: " + words[0])

		}

		if err != nil {
			fmt.Println(err.Error())
		}

		if shouldQuit {
			break
		}

	}

}
