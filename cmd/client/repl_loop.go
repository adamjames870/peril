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
			mv, errGetMove := s.gameState.CommandMove(words)
			if errGetMove != nil {
				err = errGetMove
			} else {
				errMove := publishMove(s.publishCh, s.userName, &mv)
				if errMove != nil {
					err = errMove
				} else {
					fmt.Println("move published")
				}
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
