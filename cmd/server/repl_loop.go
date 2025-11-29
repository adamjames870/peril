package main

import (
	"errors"
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
)

func ReplLoop(s *serverState) {

	for {
		words := gamelogic.GetInput()
		shouldQuit := false

		if len(words) == 0 {
			continue
		}

		var err error
		switch GetServerCommand(words[0]) {
		case Pause:
			fmt.Println("Pausing")
			err = PublishPause(s.ch)
		case Resume:
			fmt.Println("Resuming")
			err = PublishResume(s.ch)
		case Quit:
			fmt.Println("Quitting")
			shouldQuit = true
			break
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
