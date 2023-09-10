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
var count = flag.Int("count", 1, "message count to send")
var productId = flag.Int64("id", 0, "change product id")
var price = flag.Float64("price", 0, "change product price")
var title = flag.String("title", "", "change product title")
var category = flag.String("category", "", "change product category")

func main() {
	flag.Parse()

	producer, err := producer.NewProducer("localhost:9092", *topic, int32(*partition))
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	event := model.Event{
		Type: model.ProductUpdated,
		Data: model.Product{
			Id:       *productId,
			Title:    *title,
			Price:    *price,
			Category: *category,
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
