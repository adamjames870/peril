package main

import (
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishPause(ch *amqp.Channel) error {

	err := pubsub.PublishJSON(
		ch,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		return err
	}

	return nil

}

func PublishResume(ch *amqp.Channel) error {

	err := pubsub.PublishJSON(
		ch,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: false,
		},
	)
	if err != nil {
		return err
	}

	return nil

}
