package main

import (
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/consumer"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/elastic"
)

func main() {
	consumer, err := consumer.NewConsumer("product-topic", "product-elastic-group")
	if err != nil {
		panic(err)
	}

	err = consumer.Subscribe()
	if err != nil {
		panic(err)
	}

	elastic, err := elastic.NewElasticClient()
	if err != nil {
		panic(err)
	}

	err = elastic.CreateIndex()
	if err != nil {
		panic(err)
	}

	for {
		err = consumer.Poll(elastic.Insert)
		if err != nil {
			panic(err)
		}
	}
}
