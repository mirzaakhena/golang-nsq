package messagebroker

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type consumerImpl struct {
	url              string
	consumerHandlers []ConsumerHandler
}

func (c *consumerImpl) StartListening(blocking bool) {

	nsqConfig := nsq.NewConfig()

	for _, ch := range c.consumerHandlers {

		consumer, err := nsq.NewConsumer(ch.Topic, ch.Channel, nsqConfig)
		if err != nil {
			panic(err.Error())
		}
		defer consumer.Stop()

		consumer.AddHandler(&internalConsumer{handler: ch.FunctionHandler})

		if err := consumer.ConnectToNSQD(c.url); err != nil {
			panic(err.Error())
		}

	}

	if blocking {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
	}

}

func (c *consumerImpl) Handle(topic, channel string, funcHandler FunctionHandler) {

	if funcHandler == nil {
		panic("FunctionHandler must not nil")
	}

	c.consumerHandlers = append(c.consumerHandlers, ConsumerHandler{
		Topic:           topic,
		Channel:         channel,
		FunctionHandler: funcHandler,
	})
}

type internalConsumer struct {
	handler FunctionHandler
}

// HandleMessage is
func (mb *internalConsumer) HandleMessage(m *nsq.Message) error {
	mb.handler(m.Body)
	return nil
}
