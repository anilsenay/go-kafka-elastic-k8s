package main

import (
	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/search/handlers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		panic(err)
	}

	fiber := fiber.New(fiber.Config{})

	handler := handlers.NewSearchHandler(client)
	handler.SetRoutes(fiber)

	fiber.Listen(":3000")
}
