package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/anilsenay/go-kafka-elastic-k8s/replication/producer"
)

var topic = flag.String("topic", "mytopic", "topic name")
var message = flag.String("message", "HELLO", "message to send")
var count = flag.Int("count", 1, "message count to send")

func main() {
	flag.Parse()

	producer, err := producer.NewProducer("127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094", *topic)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for i := 0; i < *count; i++ {
		msg := *message + "-" + strconv.Itoa(i)
		err = producer.Produce(msg)
		if err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
