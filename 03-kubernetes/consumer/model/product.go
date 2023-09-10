package model

type Product struct {
	Id       int64   `json:"id"`
	Title    string  `json:"title,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Category string  `json:"category,omitempty"`
}
