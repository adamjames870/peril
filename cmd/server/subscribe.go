package main

import "errors"

type Sub func(clientState *serverState) error

var subscriptions = []Sub{
	subGameLogs,
}

func subscribeToQueues(state *serverState) error {

	var errs []error

	for _, subFunc := range subscriptions {
		err := subFunc(state)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)

}
