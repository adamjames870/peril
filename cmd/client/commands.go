package main

type ClientCommand int

const (
	Unknown ClientCommand = iota
	Spawn
	Move
	Status
	Help
	Spam
	Quit
)

func (c ClientCommand) String() string {
	switch c {
	case Spawn:
		return "spawn"
	case Move:
		return "move"
	case Status:
		return "status"
	case Help:
		return "help"
	case Spam:
		return "spam"
	case Quit:
		return "quit"
	default:
		return "unknown"
	}
}

func GetClientCommand(s string) ClientCommand {
	switch s {
	case "spawn":
		return Spawn
	case "move":
		return Move
	case "status":
		return Status
	case "help":
		return Help
	case "spam":
		return Spam
	case "quit":
		return Quit
	default:
		return Unknown
	}
}
