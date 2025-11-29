package main

import (
	"errors"
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
)

func ReplLoop(s *clientState) {

	for {
		words := gamelogic.GetInput()
		shouldQuit := false

		if len(words) == 0 {
			continue
		}

		var err error
		switch GetClientCommand(words[0]) {
		case Spawn:
			err = s.gameState.CommandSpawn(words)

		case Move:
			mv, mvErr := s.gameState.CommandMove(words)
			if mvErr != nil {
				err = mvErr
			} else {
				fmt.Printf("move %s %d\n", mv.ToLocation, len(mv.Units))
			}

		case Status:
			s.gameState.CommandStatus()

		case Help:
			gamelogic.PrintClientHelp()

		case Spam:
			fmt.Println("Spamming not allowed yet!")

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
