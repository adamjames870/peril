package main

import (
	"fmt"
	"os"

	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

const conn_str string = "amqp://guest:guest@localhost:5672/"

func main() {

	fmt.Println("Starting Peril server...")

	conn, errConn := amqp.Dial(conn_str)
	if errConn != nil {
		fmt.Println("Failed to load connection: " + errConn.Error())
		os.Exit(1)
	}

	fmt.Println("Opened connection to amqp")

	defer conn.Close()

	ch, errCh := conn.Channel()
	if errCh != nil {
		fmt.Println("Failed to create channel: " + errCh.Error())
		os.Exit(1)
	}

	err := pubsub.PublishJSON(
		ch,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		return
	}

	fmt.Println("Shutting down")

}
