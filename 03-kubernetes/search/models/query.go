package models

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type SearchQueryParams struct {
	Keyword  string         `query:"keyword"`
	Offset   *int           `query:"offset"`
	Limit    *int           `query:"limit"`
	MinPrice *types.Float64 `query:"min_price"`
	MaxPrice *types.Float64 `query:"max_price"`
}
