package accounting

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/imdario/mergo"
	"gorm.io/datatypes"
)

// Transaction ...
type Transaction struct {
	ID       uint       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"-"`
	DateTime time.Time  `gorm:"column:date_time;type:datetime" json:"date_time"`
	Type     OrderTypes `gorm:"column:transaction_type;" json:"type"`
}

// Customer ...
type Customer struct {
	ID         uint      `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"-"`
	CustomerID uuid.UUID `gorm:"column:customer_id;type:bigint" json:"customer_id"`
	Registered datatypes.Date
}

// ExchangeOrder ...
type ExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
}

// OrderTypes ...
type OrderTypes int64

const (
	// Withdrawal is 0
	Withdrawal OrderTypes = iota
	// Buy is 1
	Buy
	// Spread is 2
	Spread
	// Fee is 3
	Fee
)

// Value - Implementation of valuer for database/sql
func (o OrderTypes) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as bool.
	return int64(o), nil
}

// Scan - Implement the database/sql scanner interface
func (o *OrderTypes) Scan(value interface{}) error {

	// if value is nil, false
	if value == nil {
		// set the value of the pointer o to OrderTypes(1)
		*o = OrderTypes(1)
		return nil
	}

	// if this is a int64 type
	if v, ok := value.(int64); ok {
		// set the value of the pointer o to OrderTypes(v)
		*o = OrderTypes(v)
		return nil
	}

	// otherwise, return an error
	return errors.New("failed to scan OrderTypes")

}

// String ...
func (o OrderTypes) String() string {

	// declare an array of strings
	// ... operator counts how many
	// items in the array (4)
	names := [...]string{
		"withdrawal",
		"buy",
		"spread",
		"fee"}

	// → `o`: It's one of the
	// values of OrderTypes constants.
	// If the constant is withdrawal,
	// then day is 0.
	//
	// prevent panicking in case of
	// `o` is out of range of OrderTypes
	if o < Withdrawal || o > Fee {
		return "Unknown"
	}

	// return the name of a OrderType
	// constant from the names array
	// above.
	return names[o]
}

// Acc ...
type Acc interface {
	Create(c []ExchangeOrder) error
}

// Accounting ...
type Accounting struct {
}

// New ...
func New() Acc {
	return &Accounting{}
}

// Create ...
func (a *Accounting) Create(c []ExchangeOrder) error {

	// assign error, created, index, blocks, query
	var err error
	var created, blocks, index int
	var query []map[string]interface{}
	var wrap []map[string]interface{}

	index = 0
	blocks = 4

	bytes, _ := json.Marshal(&c)
	json.Unmarshal(bytes, &query)

	num := len(query)
	created = num * blocks

	result := make([]map[string]interface{}, created)

	for index < created {

		which := index % blocks
		otype := OrderTypes(which).String()

		/*NOT FINISHED*/

		result := a.rewrite(otype, result, query, index)
		mergo.Merge(wrap[index], result)

		if which == 0 {
			which++
		}

		index++

	}

	return err

}

func (a *Accounting) rewrite(o string, m []map[string]interface{}, q []map[string]interface{}, i int) []map[string]interface{} {

	for k, v := range q[i] {

		/*NOT FINISHED*/

		log.Printf("%s:%v", k, v)

		// switch k {

		// case "id":
		// 	id := uuid.New()
		// 	m[k] = id

		// case "customer_id":
		// 	m[k] = v

		// case "type":
		// 	m[k] = v

		// }

	}

	return m

}
