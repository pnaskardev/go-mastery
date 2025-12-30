package models

import "time"

type BaseEntity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Customer struct {
	BaseEntity
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	ItemID   string  `json:"item_id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Payment struct {
	Method        string `json:"method"`
	Status        string `json:"status"`
	TransactionID string `json:"transaction_id"`
}

type Order struct {
	BaseEntity
	Status      string   `json:"status"`
	TotalAmount float64  `json:"total_amount"`
	Currency    string   `json:"currency"`
	Customer    Customer `json:"customer"`
	Items       []Item   `json:"items"`
	Payment     Payment  `json:"payment"`
}
