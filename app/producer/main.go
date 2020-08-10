package main

import "github.com/mirzaakhena/belajar-nsq/messagebroker"

func main() {
	p := messagebroker.NewProducer("localhost:4150")

	err := p.Publish("Test", []byte("Hello"))
	if err != nil {
		panic(err)
	}

}
