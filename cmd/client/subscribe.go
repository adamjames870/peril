package main

import "errors"

type Sub func(clientState *clientState) error

var subscriptions = []Sub{
	subPause,
	subMoves,
}

func subscribeToQueues(state *clientState) error {

	var errs []error

	for _, subFunc := range subscriptions {
		err := subFunc(state)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)

}
