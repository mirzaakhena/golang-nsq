package main

import "github.com/mirzaakhena/belajar-nsq/messagebroker"

func main() {
	p := messagebroker.NewProducer("localhost:4150")

	if err := p.Publish("onOrderCreated", []byte("order 1234 is ready")); err != nil {
		panic(err)
	}

}
