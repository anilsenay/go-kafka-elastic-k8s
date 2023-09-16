package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/model"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	topic    string
	consumer *kafka.Consumer
}

func NewConsumer(servers, topic, groupId string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
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

func (c *Consumer) Poll(callback func(model.Event) error) error {
	ev := c.consumer.Poll(100)
	switch e := ev.(type) {
	case *kafka.Message:
		var event model.Event
		err := json.Unmarshal(e.Value, &event)
		if err != nil {
			return err
		}

		fmt.Printf("Consumed event: %s product: %d\n", event.Type, event.Data.Id)
		err = callback(event)
		if err != nil {
			return err
		}
	case kafka.Error:
		return e
	}
	return nil
}
