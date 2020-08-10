package messagebroker

// NewProducer is
func NewProducer(url string) Publisher {
	return &producer{
		defaultProducer: nsqProducerInit(url),
	}
}

// RunConsumer is
func RunConsumer(url string, consumers []ConsumerHandler) {
	nsqConsumerInit(url, consumers)
}
