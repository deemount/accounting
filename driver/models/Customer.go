package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// Customer ...
type Customer struct {
	ID         uint           `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"-"`
	CustomerID uuid.UUID      `gorm:"column:customer_id;type:bigint" json:"customer_id"`
	Attributes datatypes.JSON `gorm:"column:attributes" json:"attributes"`
	Registered datatypes.Date `gorm:"column:registered;" json:"registered"`
}
