package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/model"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/producer"
)

var topic = flag.String("topic", "product-topic", "topic name")
var partition = flag.Int("partition", 0, "partition number")
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
		product := model.GenerateProduct()
		msg, err := json.Marshal(product)
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
