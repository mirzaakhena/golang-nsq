package main

import (
	"fmt"

	"github.com/mirzaakhena/belajar-nsq/messagebroker"
)

func main() {

	c := messagebroker.NewConsumer("localhost:4150")

	c.Handle("onOrderCreated", "channel1", func(message []byte) {
		fmt.Printf("onOrderCreated >> %v\n", string(message))
	})

	c.Handle("onOrderCancelled", "channel1", func(message []byte) {
		fmt.Printf("onOrderCancelled >> %v\n", string(message))
	})

	c.StartListening(true)

}
