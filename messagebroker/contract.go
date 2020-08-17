package messagebroker

// Producer is
type Producer interface {
	Publish(topic string, message []byte) error
}

// FunctionHandler is
type FunctionHandler func(message []byte) //

// ConsumerHandler is
type ConsumerHandler struct {
	Topic           string          //
	Channel         string          //
	FunctionHandler FunctionHandler //
}

type Consumer interface {
	Handle(topic, channel string, funcHandler FunctionHandler)
	StartListening(blocking bool)
}
