package main

// SERVER

import (
	"fmt"
	"os"

	"github.com/adamjames870/peril/internal/gamelogic"
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

	errSubs := subscribeToQueues(state)
	if errSubs != nil {
		return errSubs
	}

	gamelogic.PrintServerHelp()
	ReplLoop(state)
	return nil

}
