package messagebroker

// NewProducer is
func NewProducer(url string) Producer {
	return &producerImpl{
		defaultProducer: producerImplInit(url),
	}
}

func NewConsumer(url string) Consumer {
	return &consumerImpl{
		url: url,
	}
}
