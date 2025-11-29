package main

// SERVER

import (
	"fmt"
	"os"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

const connStr string = "amqp://guest:guest@localhost:5672/"

func main() {

	fmt.Println("Starting Peril server...")

	state := serverState{
		connStr: connStr,
	}

	err := run(&state)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Shutting down")

}

func run(state *serverState) error {

	conn, errConn := amqp.Dial(state.connStr)
	if errConn != nil {
		return fmt.Errorf("failed to connect to server: %w", errConn)
	}

	state.conn = conn

	defer func() {
		if err := state.conn.Close(); err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}()

	fmt.Println("Opened connection to amqp")

	pubCh, errCh := conn.Channel()
	if errCh != nil {
		return fmt.Errorf("failed to open channel: %w", errCh)
	}

	state.publishCh = pubCh

	defer func() {
		if err := state.publishCh.Close(); err != nil {
			fmt.Println("Failed to close channel: " + err.Error())
		}
	}()

	queueName := routing.GameLogSlug
	routingKey := fmt.Sprintf("%s.*", routing.GameLogSlug)
	ch, qu, errBind := pubsub.DeclareAndBind(state.conn, routing.ExchangePerilTopic, queueName, routingKey, pubsub.Durable)
	if errBind != nil {
		return fmt.Errorf("failed to declare and bind queue: %w", errBind)
	}
	state.topicQueueName = qu.Name
	errTopicChClose := ch.Close()
	if errTopicChClose != nil {
		fmt.Println("failed to close topic declare channel:", errTopicChClose)
	}

	gamelogic.PrintServerHelp()
	ReplLoop(state)
	return nil

}
