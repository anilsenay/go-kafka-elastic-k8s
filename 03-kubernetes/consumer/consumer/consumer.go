package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	topic    string
	consumer *kafka.Consumer
}

func NewConsumer(topic, groupId string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          groupId,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
		topic:    topic,
	}, nil
}

func (c *Consumer) Subscribe() error {
	return c.consumer.Subscribe(c.topic, nil)
}

func (c *Consumer) Poll(callback func(model.Product) error) error {
	ev := c.consumer.Poll(100)
	switch e := ev.(type) {
	case *kafka.Message:
		var product model.Product
		err := json.Unmarshal(e.Value, &product)
		if err != nil {
			return err
		}
		fmt.Printf("Consumed product: %d\n", product.Id)
		err = callback(product)
		if err != nil {
			return err
		}
	case kafka.Error:
		return e
	}
	return nil
}
