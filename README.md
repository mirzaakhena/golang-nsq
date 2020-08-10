# NSQ in GO

Open 3 console

First console run the NSQ
```
$ cd docker/
$ ./run.sh 
```

Second console run Consumer
```
$ cd app/consumer/
$ go run main.go 
```
consumer will subscribe the event with topic Test

Third console run Producer
```
$ cd app/producer/
$ go run main.go 
```
producer will publish message the message with topic Test

