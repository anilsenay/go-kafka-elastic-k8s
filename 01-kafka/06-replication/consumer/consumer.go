package main

import (
	"flag"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = flag.String("topic", "mytopic", "topic name")
var groupId = flag.String("group", "mygroup", "group id")

func main() {
	flag.Parse()

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094",
		"group.id":          *groupId,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	err = consumer.Subscribe(*topic, nil)
	if err != nil {
		panic(err)
	}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Consumed message: %s\n", string(e.Value))
		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
		}
	}
}
