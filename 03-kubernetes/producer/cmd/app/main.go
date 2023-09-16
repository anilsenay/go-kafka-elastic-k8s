package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/model"
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/producer/producer"
	"github.com/gofiber/fiber/v2"
)

var topic = flag.String("topic", "product-topic", "topic name")
var partition = flag.Int("partition", 0, "partition number")

type QueryParams struct {
	Count     int   `query:"count"`
	ProductId int64 `query:"id"`
}

func main() {
	flag.Parse()

	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	producer, err := producer.NewProducer(brokers, *topic, int32(*partition))
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	app := fiber.New(fiber.Config{})

	app.Post("/", func(c *fiber.Ctx) error {
		var query QueryParams
		err := c.QueryParser(&query)
		if err != nil {
			c.SendStatus(400)
		}

		for i := 0; i < query.Count; i++ {
			event := model.Event{
				Type: model.ProductCreated,
				Data: model.GenerateProduct(),
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

		return c.SendString("OK")
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		var query QueryParams
		err := c.QueryParser(&query)
		if err != nil {
			c.SendStatus(400)
		}

		event := model.Event{
			Type: model.ProductDeleted,
			Data: model.Product{
				Id: query.ProductId,
			},
		}

		msg, err := json.Marshal(event)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
		err = producer.Produce(string(msg))
		if err != nil {
			c.Status(500).SendString(err.Error())
		}

		return c.SendString("OK")
	})

	app.Put("/", func(c *fiber.Ctx) error {
		var query model.Product
		err := c.BodyParser(&query)
		if err != nil {
			c.SendStatus(400)
		}

		event := model.Event{
			Type: model.ProductUpdated,
			Data: model.Product{
				Id:       query.Id,
				Title:    query.Title,
				Price:    query.Price,
				Category: query.Category,
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

		return c.SendString("OK")
	})

	app.Listen(":8080")
}
