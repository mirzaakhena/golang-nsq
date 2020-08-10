package main

import (
	"fmt"

	"github.com/mirzaakhena/belajar-nsq/messagebroker"
)

func main() {

	cs := []messagebroker.ConsumerHandler{
		{
			Channel: "1",
			Topic:   "Test",
			FunctionHandler: func(message []byte) error {
				fmt.Printf("1>>> %v\n", string(message))
				return nil
			},
		},
	}

	exit := make(chan int)

	messagebroker.RunConsumer("localhost:4150", cs)

	<-exit

}
