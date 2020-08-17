# NSQ in GO

## How to run this apps?
Open 3 console

In first console run the NSQ server (Of course, you need docker for running this NSQ, or you can install the NSQ manually)
```
$ cd docker/
$ ./run.sh 
```

In second console run Consumer App
```
$ cd app/consumer/
$ go run main.go 
```
Consumer will subscribing the event with topic `onOrderCreated` and `onOrderCancelled` and it will start listening the event coming from NSQ. 

In third console run Producer App
```
$ cd app/producer/
$ go run main.go 
```
Producer will publish message the message with topic `onOrderCreated`. Try to change the topic into `onOrderCancelled` and run the producer again.

## What is the channel and how do we use it?
If you have more than one consumer with the SAME TOPIC and SAME CHANNEL then ONLY ONE consumer will consume the event. You can use this scenario if you have more than one cloned consumer instance.

If you have more than one consumer with the SAME TOPIC and DIFFERENT CHANNEL then ALL the consumer will consume the event. This scenario is suitable if you have more than consumer that have different purpose but need to listen the same topic.

