package main

import amqp "github.com/rabbitmq/amqp091-go"

type serverState struct {
	connStr string
	conn    *amqp.Connection
	ch      *amqp.Channel
}
