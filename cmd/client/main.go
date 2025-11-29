package main

// CLIENT

import (
	"fmt"
	"os"

	"github.com/adamjames870/peril/internal/gamelogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

const connStr string = "amqp://guest:guest@localhost:5672/"

func main() {

	fmt.Println("Starting Peril client...")
	state := clientState{}

	err := run(&state)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Shutting down")

}

func run(state *clientState) error {

	conn, errConn := amqp.Dial(connStr)
	if errConn != nil {
		return fmt.Errorf("failed to load connection: %w", errConn)
	}
	state.conn = conn

	defer func() {
		if err := state.conn.Close(); err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}()

	fmt.Println("Opened connection to amqp")

	userName, errWelcome := gamelogic.ClientWelcome()
	if errWelcome != nil {
		fmt.Println("Failed to load welcome message: " + errWelcome.Error())
	}

	state.userName = userName
	state.gameState = gamelogic.NewGameState(userName)

	errSubs := subscribeToQueues(state)
	if errSubs != nil {
		return errSubs
	}

	gamelogic.PrintClientHelp()

	ReplLoop(state)

	return nil

}
