package models

type Product struct {
	Id         int64   `json:"id"`
	Title      string  `json:"title"`
	Price      float64 `json:"min_price"`
	CategoryId string  `json:"category_id"`
}
