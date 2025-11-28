package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/adamjames870/peril/internal/gamelogic"
	"github.com/adamjames870/peril/internal/pubsub"
	"github.com/adamjames870/peril/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

const connStr string = "amqp://guest:guest@localhost:5672/"

func main() {

	fmt.Println("Starting Peril client...")

	conn, errConn := amqp.Dial(connStr)
	if errConn != nil {
		fmt.Println("Failed to load connection: " + errConn.Error())
		os.Exit(1)
	}

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection: " + err.Error())
		}
	}(conn)

	fmt.Println("Opened connection to amqp")

	userName, errWelcome := gamelogic.ClientWelcome()
	if errWelcome != nil {
		fmt.Println("Failed to load welcome message: " + errWelcome.Error())
	}

	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, userName)

	_, _, errDecBnd := pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, queueName, routing.PauseKey, pubsub.Transient)
	if errDecBnd != nil {
		fmt.Println("Failed to publish and bind: " + errDecBnd.Error())
		os.Exit(1)
	}

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

}
