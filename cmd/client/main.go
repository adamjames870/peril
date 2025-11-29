package main

// CLIENT

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

	fmt.Println("Starting Peril client...")
	state := clientState{}

	conn, errConn := amqp.Dial(connStr)
	if errConn != nil {
		fmt.Println("Failed to load connection: " + errConn.Error())
		os.Exit(1)
	}
	state.conn = conn

	defer func(conn *amqp.Connection) {
		err := state.conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection: " + err.Error())
		}
	}(state.conn)

	fmt.Println("Opened connection to amqp")

	userName, errWelcome := gamelogic.ClientWelcome()
	if errWelcome != nil {
		fmt.Println("Failed to load welcome message: " + errWelcome.Error())
	}
	state.userName = userName

	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, userName)

	ch, qu, errDecBnd := pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, queueName, routing.PauseKey, pubsub.Transient)
	if errDecBnd != nil {
		fmt.Println("Failed to publish and bind: " + errDecBnd.Error())
		os.Exit(1)
	}
	state.ch = ch
	state.qu = &qu

	state.gameState = gamelogic.NewGameState(userName)
	gamelogic.PrintClientHelp()

	ReplLoop(state)

}
