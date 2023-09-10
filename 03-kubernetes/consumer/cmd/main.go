package main

import (
	"fmt"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/consumer"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/elastic"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/model"
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
		err = consumer.Poll(func(e model.Event) error {
			switch e.Type {
			case model.ProductCreated:
				return elastic.Insert(e.Data)
			case model.ProductDeleted:
				return elastic.Delete(e.Data)
			case model.ProductUpdated:
				return elastic.Update(e.Data)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	}
}
