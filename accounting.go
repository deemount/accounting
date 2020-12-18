package accounting

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/imdario/mergo"

	"github.com/deemount/accounting/asserts"
	"github.com/deemount/accounting/driver"
)

// ExchangeOrder ...
type ExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
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
	db := driver.NewDataService(*driver.DataService.Config)
	idle, err := db.Connect()
	if err != nil {
		log.Printf("Could not open database connection: %v", err)
	}
	log.Print(idle)
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
		otype := asserts.OrderTypes(which).String()

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
