package main

type ServerCommand int

const (
	Unknown ServerCommand = iota
	Pause
	Resume
	Quit
)

func (c ServerCommand) String() string {
	switch c {
	case Pause:
		return "pause"
	case Resume:
		return "resume"
	case Quit:
		return "quit"
	default:
		return "unknown"
	}
}

func GetServerCommand(s string) ServerCommand {
	switch s {
	case "pause":
		return Pause
	case "resume":
		return Resume
	case "quit":
		return Quit
	default:
		return Unknown
	}
}
