package handlers

import (
	"context"
	"encoding/json"

	"github.com/anilsenay/go-kafka-elastic-k8s/elasticsearch/search/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gofiber/fiber/v2"
)

type SearchHandler struct {
	elastic *elasticsearch.TypedClient
}

func NewSearchHandler(e *elasticsearch.TypedClient) *SearchHandler {
	return &SearchHandler{
		elastic: e,
	}
}

func (h *SearchHandler) handleSearchById(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.elastic.Get("cimri_product_index", id).Do(context.TODO())
	if !res.Found {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var product models.Product
	json.Unmarshal(res.Source_, &product)

	return c.JSON(product)
}

func (h *SearchHandler) handleSearch(c *fiber.Ctx) error {
	var query models.SearchQueryParams
	err := c.QueryParser(&query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.elastic.Search().
		Index("cimri_product_index").
		Request(&search.Request{
			From: query.Offset,
			Size: query.Limit,
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"title": {
									Query: query.Keyword,
								},
							},
						},
						{
							Range: map[string]types.RangeQuery{
								"price": types.NumberRangeQuery{
									Gte: query.MinPrice,
									Lte: query.MaxPrice,
								},
							},
						},
					},
				},
			},
		}).
		Do(context.TODO())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(res.Hits.Hits) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	var products = make([]models.Product, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		var product models.Product
		json.Unmarshal(hit.Source_, &product)
		products[i] = product
	}

	return c.JSON(products)
}

func (h *SearchHandler) SetRoutes(app *fiber.App) {
	app.Get("/search/:id", h.handleSearchById)
	app.Get("/search", h.handleSearch)
}
