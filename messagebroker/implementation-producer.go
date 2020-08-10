package messagebroker

import "github.com/nsqio/go-nsq"

// producer is
type producer struct {
	defaultProducer *nsq.Producer
}

func nsqProducerInit(url string) *nsq.Producer {
	nsqConfig := nsq.NewConfig()
	nsqProducer, err := nsq.NewProducer(url, nsqConfig)
	if err != nil {
		panic(err.Error())
	}
	return nsqProducer
}

func (x *producer) Publish(topic string, body []byte) error {
	return x.defaultProducer.Publish(topic, body)
}
