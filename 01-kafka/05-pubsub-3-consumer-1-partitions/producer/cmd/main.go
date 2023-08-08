package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/anilsenay/go-kafka-elastic-k8s/pubsub-3-consumer-1-partitions/producer"
)

var topic = flag.String("topic", "mytopic", "topic name")
var partition = flag.Int("partition", 0, "partition number")
var message = flag.String("message", "HELLO", "message to send")
var count = flag.Int("count", 1, "message count to send")

func main() {
	flag.Parse()

	producer, err := producer.NewProducer("localhost:9092", *topic, int32(*partition))
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	fmt.Println(*count)
	for i := 0; i < *count; i++ {
		msg := *message + "-" + strconv.Itoa(i)
		err = producer.Produce(msg)
		if err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
