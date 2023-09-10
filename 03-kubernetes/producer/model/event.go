package model

type EventTypes string

const (
	ProductCreated EventTypes = "product_created"
	ProductUpdated EventTypes = "product_updated"
	ProductDeleted EventTypes = "product_deleted"
)

type Event struct {
	Type EventTypes `json:"type"`
	Data Product    `json:"data"`
}
