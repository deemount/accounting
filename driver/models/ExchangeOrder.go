package models

import "github.com/google/uuid"

// ExchangeOrder ...
type ExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
}
