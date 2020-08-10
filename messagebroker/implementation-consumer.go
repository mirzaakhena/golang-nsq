package messagebroker

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

var consumers []*nsq.Consumer

func nsqConsumerInit(url string, consumerHandlers []ConsumerHandler) {

	consumers = []*nsq.Consumer{}

	nsqConfig := nsq.NewConfig()

	// wg := &sync.WaitGroup{}
	for _, ch := range consumerHandlers {

		if ch.FunctionHandler == nil {
			panic("FunctionHandler must not nil")
		}

		consumer, err := nsq.NewConsumer(ch.Topic, ch.Channel, nsqConfig)
		if err != nil {
			panic(err.Error())
		}

		consumer.AddHandler(&internalConsumer{handler: ch.FunctionHandler})

		if err := consumer.ConnectToNSQD(url); err != nil {
			panic(err.Error())
		}
		// wg.Add(1)

		fmt.Printf("topic: %v\n", ch.Topic)
	}

	// wg.Wait()

}

type internalConsumer struct {
	handler FunctionHandler
}

// HandleMessage is
func (mb *internalConsumer) HandleMessage(m *nsq.Message) error {
	if mb.handler == nil {
		return nil
	}
	return mb.handler(m.Body)
}
