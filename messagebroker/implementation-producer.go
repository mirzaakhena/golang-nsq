package messagebroker

import "github.com/nsqio/go-nsq"

// producerImpl is
type producerImpl struct {
	defaultProducer *nsq.Producer
}

func producerImplInit(url string) *nsq.Producer {
	nsqConfig := nsq.NewConfig()
	nsqProducer, err := nsq.NewProducer(url, nsqConfig)
	if err != nil {
		panic(err.Error())
	}
	return nsqProducer
}

func (p *producerImpl) Publish(topic string, body []byte) error {
	return p.defaultProducer.Publish(topic, body)
}
