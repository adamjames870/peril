package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/gamelogic"
)

func handler_move(gs *gamelogic.GameState) func(move gamelogic.ArmyMove) {

	return func(move gamelogic.ArmyMove) {
		defer fmt.Print("> ")
		gs.HandleMove(move)
	}

}
