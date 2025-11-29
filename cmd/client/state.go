package main

import (
	"github.com/adamjames870/peril/internal/gamelogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

type clientState struct {
	connStr   string
	conn      *amqp.Connection
	userName  string
	ch        *amqp.Channel
	qu        *amqp.Queue
	gameState *gamelogic.GameState
}
