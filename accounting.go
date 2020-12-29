package accounting

import (
	"encoding/json"
	"log"

	"github.com/imdario/mergo"

	"github.com/deemount/accounting/app"
	"github.com/deemount/accounting/asserts"
	"github.com/deemount/accounting/driver"
	"github.com/deemount/accounting/driver/models"
)

// Service is a struct
type Service struct {
	App app.App
}

var service = Service{}

// Acc ...
type Acc interface {
	Create(c []models.ExchangeOrder) error
}

// Accounting ...
type Accounting struct {
}

// New ...
func New() Acc {
	return &Accounting{}
}

// Init ...
func (acc *Accounting) Init() {

	db := driver.NewDataService(*service.App.DB.Config)
	idle, err := db.Connect()
	if err != nil {
		log.Printf("Could not open database connection: %v", err)
	}
	service.App.DB = idle

}

// Create ...
func (acc *Accounting) Create(c []models.ExchangeOrder) error {

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

		result := acc.rewrite(otype, result, query, index)
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
