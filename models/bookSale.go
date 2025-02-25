package models

type BookSale struct {
	Book     `json:"book"`
	Quantity int `json:"quantity_sold"`
}
