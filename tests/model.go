package tests

import "github.com/google/uuid"

// TestExchangeOrder ...
type TestExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
	Asset      string    `json:"asset"`
}
