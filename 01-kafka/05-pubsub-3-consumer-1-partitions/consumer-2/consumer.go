package main

import (
	"flag"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = flag.String("topic", "mytopic", "topic name")
var partition = flag.Int("partition", -1, "partition number")
var groupId = flag.String("group", "mygroup2", "group id")

func main() {
	flag.Parse()

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          *groupId,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	if *partition == -1 {
		err = consumer.Subscribe(*topic, nil)
	} else {
		err = consumer.Assign([]kafka.TopicPartition{
			{Topic: topic, Partition: int32(*partition)},
		})
	}

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
