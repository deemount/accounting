package asserts

import (
	"database/sql/driver"
	"errors"
)

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
