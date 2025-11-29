package main

import "errors"

type Sub func(clientState *clientState) error

func GetSubscriptions() []Sub {
	return []Sub{
		subPause,
		subMoves,
	}
}

func subscribeToQueues(state *clientState) error {

	var errs []error

	for _, subFunc := range GetSubscriptions() {
		err := subFunc(state)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)

}
