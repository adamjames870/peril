package main

// SERVER

import (
	"errors"
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

	conn, errConn := amqp.Dial(state.connStr)
	if errConn != nil {
		fmt.Println("Failed to load connection: " + errConn.Error())
		os.Exit(1)
	}

	state.conn = conn

	fmt.Println("Opened connection to amqp")

	defer func(conn *amqp.Connection) {
		err := state.conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection: " + err.Error())
		}
	}(state.conn)

	ch, errCh := conn.Channel()
	if errCh != nil {
		fmt.Println("Failed to create channel: " + errCh.Error())
		os.Exit(1)
	}
	state.ch = ch

	gamelogic.PrintServerHelp()
	ReplLoop(&state)

	fmt.Println("Shutting down")

}
