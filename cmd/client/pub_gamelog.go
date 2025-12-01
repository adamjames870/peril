package main

import (
	"fmt"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
)

func publishGameLog(s *clientState, log routing.GameLog) error {

	routingKey := fmt.Sprintf("%s.%s", routing.GameLogSlug, log.Username)

	return pubsub.PublishGob(
		s.publishCh,
		routing.ExchangePerilTopic,
		routingKey,
		log,
	)

}
