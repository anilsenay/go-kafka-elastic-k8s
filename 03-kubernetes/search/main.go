package main

import (
	"os"
	"strings"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/search/handlers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	es_user := os.Getenv("ELASTICSEARCH_USER")
	es_pass := os.Getenv("ELASTICSEARCH_PASSWORD")

	es_hosts := os.Getenv("ELASTICSEARCH_HOSTS")
	if es_hosts == "" {
		es_hosts = "localhost:9200"
	}

	addresses := strings.Split(es_hosts, ",")

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: addresses,
		Username:  es_user,
		Password:  es_pass,
	})
	if err != nil {
		panic(err)
	}

	fiber := fiber.New(fiber.Config{})

	handler := handlers.NewSearchHandler(client)
	handler.SetRoutes(fiber)

	fiber.Listen(":3000")
}
