package elastic

import (
	"context"
	"fmt"

	"github.com/anilsenay/go-kafka-elastic-k8s/kubernetes/consumer/model"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type ElasticClient struct {
	client *elasticsearch8.TypedClient
}

func NewElasticClient() (*ElasticClient, error) {
	client, err := elasticsearch8.NewTypedClient(elasticsearch8.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		return nil, err
	}
	return &ElasticClient{
		client: client,
	}, nil
}

func (e *ElasticClient) CreateIndex() error {
	existReq := e.client.Indices.Exists("cimri_product_index")
	exist, err := existReq.Do(context.Background())
	if err != nil {
		return err
	}
	if !exist {
		_, err := e.client.Indices.Create("cimri_product_index").Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ElasticClient) Insert(product model.Product) error {
	resp, err := e.client.Index("cimri_product_index").
		Id(fmt.Sprintf("%d", product.Id)).
		Request(product).
		Do(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("Product with id: %s created\n", resp.Id_)
	return nil
}
