package models

import (
	"time"

	"github.com/deemount/accounting/asserts"
)

// Transaction ...
type Transaction struct {
	ID       uint               `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"-"`
	DateTime time.Time          `gorm:"column:date_time;type:datetime" json:"date_time"`
	Type     asserts.OrderTypes `gorm:"column:transaction_type;" json:"type"`
}
