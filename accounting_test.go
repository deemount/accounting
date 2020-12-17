package accounting

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// ExchangeOrder ...
type TestExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
}

type TestOrderTypes int64

const (
	TestWithdrawal TestOrderTypes = iota
	TestBuy
	TestSpread
	TestFee
)

// String ...
func (testo TestOrderTypes) String() string {

	names := [...]string{
		"withdrawal",
		"buy",
		"spread",
		"fee"}

	if testo < TestWithdrawal || testo > TestFee {
		return "Unknown"
	}

	return names[testo]
}

// TestOrderTypesValues
// cmd: go test -v accounting_test.go -run TestOrderTypesValues
func TestOrderTypesValues(t *testing.T) {

	//
	t.Run("Results", func(t *testing.T) {
		t.Logf("Withdrawal: %v\n", TestOrderTypes(0))
		t.Logf("Buy: %v\n", TestOrderTypes(1))
		t.Logf("Spread: %v\n", TestOrderTypes(2))
		t.Logf("Fee: %v\n", TestOrderTypes(3))
	})
	t.Log("\n\n")

	//
	t.Run("TypeOf", func(t *testing.T) {

		// rt1 := reflect.TypeOf(TestOrderTypes(1))
		// rt2 := reflect.TypeOf(int64(1))

		// assert.Equal(t, rt1, rt2, "is not equal, because TestOrderTypes is a unsigned pointer and differs to unsigned integer")

	})
	t.Log("\n\n")

	//
	t.Run("PointerTo", func(t *testing.T) {

		rt3 := reflect.TypeOf(TestOrderTypes(0))
		rt4 := reflect.TypeOf(new(TestOrderTypes))

		assert.Equal(t, reflect.PtrTo(rt3), rt4, fmt.Sprintf("%s", "is equal, because both have same type of pointer"))

	})
	t.Log("\n\n")

	//
	t.Run("ValueOf", func(t *testing.T) {

		// rt5 := reflect.ValueOf(TestOrderTypes(1))
		// rt6 := reflect.ValueOf(int64(1))

		// assert.Equal(t, rt5, rt6, "is not equal")
		// assert.Equal(t, reflect.Kind(TestOrderTypes(1)), reflect.Kind(int64(1)), "is equal")

	})

}

// TestStruct2Map
// cmd: go test -v accounting_test.go -run TestStruct2Map
func TestStruct2Map(t *testing.T) {

	//
	t.Run("Single", func(t *testing.T) {

		q1 := make(map[string]interface{}, 1)

		c1 := &TestExchangeOrder{
			ID:         uuid.New(),
			CustomerID: 1,
			Type:       "buy",
		}

		bytes, _ := json.Marshal(&c1)
		json.Unmarshal(bytes, &q1)

		t.Logf("Single: %+v", q1)

	})
	t.Log("\n\n")

	//
	t.Run("Many", func(t *testing.T) {

		q2 := make([]map[string]interface{}, 2)

		c2 := &[]TestExchangeOrder{
			{
				ID:         uuid.New(),
				CustomerID: 1,
				Type:       "buy",
			},
			{
				ID:         uuid.New(),
				CustomerID: 2,
				Type:       "buy",
			},
		}

		bytes, _ := json.Marshal(&c2)
		json.Unmarshal(bytes, &q2)

		t.Logf("Many: %v", q2)

	})

}

// TestMakeMaps
// cmd: go test -v accounting_test.go -run TestMakeMaps
func TestMakeMaps(t *testing.T) {

	t.Run("Many", func(t *testing.T) {

		num := 4
		emptyMaps := make([]map[string]interface{}, num)
		t.Logf("%v", emptyMaps)

	})

}

// TestLoopMap
// cmd: go test -v accounting_test.go -run TestLoopMap
func TestLoopMap(t *testing.T) {

	t.Run("Slice", func(t *testing.T) {

		var created, blocks, index int
		var query []map[string]interface{}
		//var wrap []map[string]interface{}

		index = 0
		blocks = 4

		var c = &[]TestExchangeOrder{
			{
				ID:         uuid.New(),
				CustomerID: 1,
				Type:       "buy",
			},
			{
				ID:         uuid.New(),
				CustomerID: 2,
				Type:       "buy",
			},
		}

		bytes, _ := json.Marshal(&c)
		json.Unmarshal(bytes, &query)

		rows := len(query)
		created = rows * blocks
		t.Logf("%v", query)

		result := make([]map[string]interface{}, created)

		t.Log("index:\nis the position of current map, starting at 0\n")
		t.Log("blocks:\nis number of wanted blocks for each map in map slice\n")
		t.Log("created:\nis number of all maps by multiplication of rows and blocks\n")
		t.Log("otype:\nis string value\n")
		t.Log("which:\nis integer value calculated with modulo and represents\nthe value of a typed constant\n")
		for index < created {

			which := index % blocks
			otype := TestOrderTypes(which).String()

			//t.Logf("%d (of %d) mod %d = %s(%d)", index, created, blocks, otype, which)

			MockRewrite(otype, result, query, index)

			if which == 0 {
				which++
			}

			index++

		}

	})

}

//
func MockRewrite(o string, r []map[string]interface{}, q []map[string]interface{}, i int) {

	log.Printf("%v", r)

	for k, v := range r {

		log.Printf("INDEX(%d):OPTYPE(%s):KEY(%d):MAP(%v)\n", i, o, k, v)

		// for str, val := range v {

		// 	switch str {
		// 	case "id":
		// 		log.Printf("-- ID: %s:%s:%v\n", o, str, val)
		// 	case "customer_id":
		// 		log.Printf("-- CUSTOMERID: %s:%s:%v\n", o, str, val)
		// 	case "type":
		// 		log.Printf("-- TYPE: %s:%s:%v\n", o, str, val)
		// 	default:
		// 		log.Print("NONE")
		// 	}

		//}

		log.Printf("\n")

	}

}

func MockRewrite2(o string, q []map[string]interface{}, i int) {

	for k, v := range q[i] {

		log.Printf("%s:%s:%v", o, k, v)

	}

}

// TestMakeMaps
// cmd: go test -v accounting_test.go -run TestForLoop
func TestForLoop(t *testing.T) {

	t.Run("Results", func(t *testing.T) {

		rows := 3
		index := 0
		blocks := 4
		created := rows * blocks

		for index < created {

			which := index % blocks
			t.Logf("%d:%s", index, TestOrderTypes(which).String())

			if which == 0 {
				which++
			}

			index++

		}

	})

}

// getSize
// This obtains the size of v using reflect and then, for the supported types
// in this example (slices, maps, strings, and structs), it computes the memory
// required by the content stored in them. You would need to add here other types
// that you need to support.
//
// There are a few details to work out:
// - Private fields are not counted.
// - For structs we are double-counting the basic types.
//
// For number two, you can filter them out before doing the recursive call
// when handling structs, you can check the kinds in the documentation for the reflect package.
func getSize(v interface{}) int {

	size := int(reflect.TypeOf(v).Size())

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			size += getSize(s.Index(i).Interface())
		}

	case reflect.Map:

		s := reflect.ValueOf(v)
		keys := s.MapKeys()
		size += int(float64(len(keys)) * 10.79) // approximation from https://golang.org/src/runtime/hashmap.go
		for i := range keys {
			size += getSize(keys[i].Interface()) + getSize(s.MapIndex(keys[i]).Interface())
		}
	case reflect.String:
		size += reflect.ValueOf(v).Len()

	case reflect.Struct:
		s := reflect.ValueOf(v)
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).CanInterface() {
				size += getSize(s.Field(i).Interface())
			}
		}

	}
	return size
}
