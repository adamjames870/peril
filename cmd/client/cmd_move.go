package main

import "fmt"

func cmdMove(state *clientState, words []string) error {

	var err error
	mv, errGetMove := state.gameState.CommandMove(words)
	if errGetMove != nil {
		err = errGetMove
	} else {
		errMove := publishMove(state.publishCh, state.userName, &mv)
		if errMove != nil {
			err = errMove
		} else {
			fmt.Println("move published")
		}
	}
	return err

}
