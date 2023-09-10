package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/model"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/producer"
)

var topic = flag.String("topic", "product-topic", "topic name")
var partition = flag.Int("partition", 0, "partition number")
var productId = flag.Int64("id", 0, "change product id")

func main() {
	flag.Parse()

	producer, err := producer.NewProducer("localhost:9092", *topic, int32(*partition))
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	event := model.Event{
		Type: model.ProductDeleted,
		Data: model.Product{
			Id: *productId,
		},
	}

	msg, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	err = producer.Produce(string(msg))
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)
}
