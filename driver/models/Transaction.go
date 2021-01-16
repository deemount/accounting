package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Transaction ...
type Transaction struct {
	TransactionID       uint            `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"-"`
	TransactionDateTime time.Time       `gorm:"column:date_time;type:datetime" json:"date_time"`
	CustomerID          int64           `json:"customerID"`
	Type                string          `json:"type"`
	Asset               string          `json:"asset"`
	Currency            string          `json:"currency"`
	Deposit             decimal.Decimal `json:"deposit"`
}
