package main

import (
	"context"
	"fmt"

	"github.com/anilsenay/go-kafka-elastic-k8s/elasticsearch/feed/random"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type Product struct {
	Id         int64   `json:"id"`
	Title      string  `json:"title"`
	MinPrice   float64 `json:"min_price"`
	CategoryId string  `json:"category_id"`
}

var categoryCache = map[int64]string{130: "Cep Telefonları"}

func main() {
	client, err := elasticsearch8.NewTypedClient(elasticsearch8.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		panic(err)
	}

	existReq := client.Indices.Exists("cimri_product_index")
	exist, err := existReq.Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exist {
		client.Indices.Create("cimri_product_index").Do(context.Background())
	}

	products := []Product{
		{Id: random.RandomId(), Title: random.GetRandomTitle(), MinPrice: random.RandomPrice(), CategoryId: categoryCache[130]},
		{Id: random.RandomId(), Title: random.GetRandomTitle(), MinPrice: random.RandomPrice(), CategoryId: categoryCache[130]},
		{Id: random.RandomId(), Title: random.GetRandomTitle(), MinPrice: random.RandomPrice(), CategoryId: categoryCache[130]},
		{Id: random.RandomId(), Title: random.GetRandomTitle(), MinPrice: random.RandomPrice(), CategoryId: categoryCache[130]},
		{Id: random.RandomId(), Title: random.GetRandomTitle(), MinPrice: random.RandomPrice(), CategoryId: categoryCache[130]},
	}

	for _, product := range products {
		resp, err := client.Index("cimri_product_index").
			Id(fmt.Sprintf("%d", product.Id)).
			Request(product).
			Do(context.Background())

		if err != nil {
			fmt.Printf("Error getting response: %s", err)
		}

		fmt.Printf("Product with id: %s created\n", resp.Id_)
	}
}