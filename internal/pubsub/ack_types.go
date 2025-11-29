package pubsub

type AckType int

const (
	Ack AckType = iota
	NackRequeue
	NackDiscard
)
