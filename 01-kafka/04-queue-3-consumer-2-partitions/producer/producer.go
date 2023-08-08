package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MyProducer struct {
	*kafka.Producer
	Topic string
	// Partition     int32
	delivery_chan chan kafka.Event
}

func NewProducer(hosts, topic string, partition int32) (*MyProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": hosts,
	})
	if err != nil {
		return nil, err
	}
	return &MyProducer{
		Producer: p,
		Topic:    topic,
		// Partition:     partition,
		delivery_chan: make(chan kafka.Event, 10000),
	}, nil
}

func (p *MyProducer) Produce(message string) error {
	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.Topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message)},
		p.delivery_chan,
	)
	if err != nil {
		return err
	}

	e := <-p.delivery_chan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	} else {
		fmt.Printf("Produced message: %s to topic: %s, partition: %d, offset: %d\n", string(m.Value), *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		return nil
	}
}

func (p *MyProducer) Close() {
	p.Producer.Close()
}
