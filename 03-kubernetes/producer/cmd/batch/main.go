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

var batch = []model.Event{
	{Type: model.ProductCreated, Data: model.Product{Id: 1000}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1001}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1002}},
	{Type: model.ProductUpdated, Data: model.Product{Id: 1000, Price: 10000}},
	{Type: model.ProductUpdated, Data: model.Product{Id: 1002, Price: 12000}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1003}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1004}},
	{Type: model.ProductUpdated, Data: model.Product{Id: 1001, Price: 11000}},
	{Type: model.ProductUpdated, Data: model.Product{Id: 1003, Title: model.GetRandomTitle()}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1005}},
	{Type: model.ProductCreated, Data: model.Product{Id: 1006}},
	{Type: model.ProductUpdated, Data: model.Product{Id: 1004, Price: 14000}},
	{Type: model.ProductDeleted, Data: model.Product{Id: 1002}},
	{Type: model.ProductDeleted, Data: model.Product{Id: 1005}},
}

func main() {
	flag.Parse()

	producer, err := producer.NewProducer("localhost:9092", *topic, int32(*partition))
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for _, b := range batch {
		event := b

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
}
