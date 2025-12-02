package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/routing"
)

func cmdSpam(state *clientState, words []string) error {
	if len(words) < 2 {
		return errors.New("not enough arguments")
	}
	num, err := strconv.Atoi(words[1])
	if err != nil {
		return errors.New("not a number")
	}

	var errs []error
	for i := 0; i < num; i++ {
		log := routing.GameLog{
			CurrentTime: time.Now(),
			Message:     gamelogic.GetMaliciousLog(),
			Username:    state.gameState.GetUsername(),
		}
		errs = append(errs, publishGameLog(state, log))
	}

	return errors.Join(errs...)
}
