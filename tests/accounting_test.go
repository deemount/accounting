package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

/*
  Test management, I
  modeling data
*/

// TestExchangeOrder represents the model of a reporting list
type TestExchangeOrder struct {
	TransactionID       uuid.UUID       `json:"transactionID"`
	TransactionDateTime time.Time       `json:"transactionDateTime"`
	CustomerID          int64           `json:"customerID"`
	Type                string          `json:"type"`
	Asset               string          `json:"asset"`
	Currency            string          `json:"currency"`
	Deposit             decimal.Decimal `json:"deposit"`
}

/*
  Test management, II
  implementing different kind of test data
*/

// untyped constant
const (
	testLayoutISO = time.RFC3339
	// testFixedNumberOfOrderTypes is also the number of wanted
	// maps for each query result
	testFixedNumberOfOrderTypes = 4
)

// assigned variables
var (

	// create date time format for testing
	testNow         = time.Now()
	testFormat      = testNow.Format(testLayoutISO)
	testDateTime, _ = time.Parse(testLayoutISO, testFormat)

	// TestQueryMap
	TestQueryMap []map[string]interface{}

	// TestMapExchangeOrder
	TestMapExchangeOrder []map[string]TestExchangeOrder

	// TestDataStruct is a filled single struct type for testing
	TestDataStruct = TestExchangeOrder{
		TransactionID: uuid.New(),
		CustomerID:    1,
		Type:          "buy",
		Asset:         "EUR",
	}

	// TestDataStructSlice is a filled struct slice for testing
	TestDataStructSlice = []TestExchangeOrder{
		{
			TransactionID:       uuid.New(),
			TransactionDateTime: testDateTime,
			CustomerID:          1,
			Type:                "buy",
			Asset:               "BTC",
			Currency:            "EUR",
		},
		{
			TransactionID:       uuid.New(),
			TransactionDateTime: testDateTime,
			CustomerID:          2,
			Type:                "buy",
			Asset:               "BTC",
			Currency:            "EUR",
		},
		{
			TransactionID:       uuid.New(),
			TransactionDateTime: testDateTime,
			CustomerID:          3,
			Type:                "buy",
			Asset:               "BTC",
			Currency:            "EUR",
		},
		{
			TransactionID:       uuid.New(),
			TransactionDateTime: testDateTime,
			CustomerID:          4,
			Type:                "buy",
			Asset:               "BTC",
			Currency:            "EUR",
		},
	}

	// TestNumberOfSlicesInTestStruct
	TestNumberOfSlicesInTestStruct = len(TestDataStructSlice)

	// TestCreatedSlices is number of all maps by multiplication
	// of TestNumberOfSlicesInTestStruct and testFixedNumberOfOrderTypes
	TestCreatedSlices = TestNumberOfSlicesInTestStruct * testFixedNumberOfOrderTypes
)

/*
  Test management, III
  create common type
*/

// TestOrderTypes ...
type TestOrderTypes int64

const (
	TestBuy TestOrderTypes = iota
	TestWithdrawal
	TestSpread
	TestFee
)

// String return name of typed constants by given integer (0,1,2,3,4)
func (ot TestOrderTypes) String() string {

	names := [...]string{
		"buy",
		"spread",
		"fee",
		"withdrawal",
	}

	if ot < TestBuy || ot > TestWithdrawal {
		return "Unknown"
	}

	return names[ot]
}

/*
  Test management, IV
  testing functionalities
*/

// TestOrderTypesValues
// cmd: go test -v accounting_test.go -run TestOrderTypesValues
func TestOrderTypesValues(t *testing.T) {

	// show integer values of types
	t.Run("Results", func(t *testing.T) {
		t.Logf("Buy: %v\n", TestOrderTypes(0))
		t.Logf("Spread: %v\n", TestOrderTypes(1))
		t.Logf("Fee: %v\n", TestOrderTypes(2))
		t.Logf("Withdrawal: %v\n", TestOrderTypes(3))
	})
	t.Log("\n\n")

	// reflect pointer of types through different initialization
	t.Run("PointerTo", func(t *testing.T) {

		rt3 := reflect.TypeOf(TestOrderTypes(0))
		rt4 := reflect.TypeOf(new(TestOrderTypes))

		ok := assert.Equal(t, reflect.PtrTo(rt3), rt4)
		if ok {
			t.Log("is equal, because both have same type of pointer")
		}
	})
	t.Log("\n\n")

	// reflect kind of order types
	t.Run("Kind", func(t *testing.T) {

		rt7 := reflect.Kind(TestOrderTypes(1))
		rt8 := reflect.Kind(int64(1))

		ok2 := assert.Equal(t, rt7, rt8)
		if ok2 {
			t.Log("reflect.Kind(TestOrderTypes(1)) and reflect.Kind(int64(1)) is equal")
		}

	})

}

// TestStruct2Map
// cmd: go test -v accounting_test.go -run TestStruct2Map
func TestStruct2Map(t *testing.T) {

	//
	t.Run("Single", func(t *testing.T) {

		q1 := make(map[string]interface{}, 1)

		bytes, _ := json.Marshal(&TestDataStruct)
		json.Unmarshal(bytes, &q1)

		t.Logf("Single: %+v", q1)

	})
	t.Log("\n\n")

	//
	t.Run("Many", func(t *testing.T) {

		q2 := make([]map[string]interface{}, TestNumberOfSlicesInTestStruct)

		bytes, _ := json.Marshal(&TestDataStructSlice)
		json.Unmarshal(bytes, &q2)

		t.Logf("Many: %v", q2)

	})

}

// TestMakeMaps
// cmd: go test -v accounting_test.go -run TestMakeMaps
func TestMakeMaps(t *testing.T) {

	t.Run("Many", func(t *testing.T) {

		emptyMaps := make([]map[string]interface{}, TestNumberOfSlicesInTestStruct)
		t.Logf("%v", emptyMaps)

	})

}

// TestAppendStructSlice
// cmd: go test -v accounting_test.go -run TestAppendStructSlice
func TestAppendStructSlice(t *testing.T) {

	t.Run("Show", func(t *testing.T) {

		// is the position of current map, starting at 0
		TestZeroIndexQueryMap := 0
		result := make([]TestExchangeOrder, TestCreatedSlices)
		for i := 0; i < TestCreatedSlices; i++ {
			if i%4 == 0 {
				result[i] = TestDataStructSlice[TestZeroIndexQueryMap]
				TestZeroIndexQueryMap++
			}
		}

		ShowSize(result)
		t.Logf("%d:%d", len(result), cap(result))
		b, _ := json.MarshalIndent(result, " ", "   ")
		t.Logf("%v", string(b))

	})

}

// TestAppendMapSlice
// cmd: go test -v accounting_test.go -run TestAppendMapSlice
func TestAppendMapSlice(t *testing.T) {

	t.Run("Show", func(t *testing.T) {

		b, _ := json.Marshal(&TestDataStructSlice)
		json.Unmarshal(b, &TestQueryMap)

		// is the position of current map, starting at 0
		TestZeroIndexQueryMap := 0
		result := make([]map[string]interface{}, TestCreatedSlices)
		for i := 0; i < TestCreatedSlices; i++ {
			if i%testFixedNumberOfOrderTypes == 0 {
				result[i] = TestQueryMap[TestZeroIndexQueryMap]
				TestZeroIndexQueryMap++
			}
		}

		ShowSize(result)
		t.Logf("%d:%d", len(result), cap(result))
		b1, _ := json.MarshalIndent(result, " ", "   ")
		t.Logf("%v", string(b1))

	})

}

// TestAppendRewriteMapSlice
// cmd: go test -v ./tests/accounting_test.go -run TestAppendRewriteMapSlice
func TestAppendRewriteMapSlice(t *testing.T) {

	t.Run("Slice", func(t *testing.T) {

		b, _ := json.Marshal(&TestDataStructSlice)
		json.Unmarshal(b, &TestQueryMap)

		// is the position of current map, starting at 0
		TestZeroIndexQueryMap := 0

		result := make([]map[string]interface{}, TestCreatedSlices)

		//
		for i := 0; i < TestCreatedSlices; i++ {
			which := i % testFixedNumberOfOrderTypes
			otype := TestOrderTypes(which).String()
			if which == 0 {
				result[i] = TestQueryMap[TestZeroIndexQueryMap]
				result[i]["type"] = otype
				TestZeroIndexQueryMap++
			} else {
				result[i] = map[string]interface{}{
					"customer_id": TestQueryMap[TestZeroIndexQueryMap-1]["customer_id"],
					"id":          TestQueryMap[TestZeroIndexQueryMap-1]["id"],
					"type":        otype,
					"asset":       "EUR",
				}
			}
			t.Logf("%d of %d (l:%d)", i, TestCreatedSlices, len(result[i]))
			ShowSize(result[i])
		}

	})

}

// TestCreate
// cmd: go test -v ./tests/accounting_test.go -run TestCreate
func TestCreate(t *testing.T) {

	b, _ := json.Marshal(&TestDataStructSlice)
	json.Unmarshal(b, &TestQueryMap)

	// is the position of query map, starting at 0
	TestZeroIndexQueryMap := 0

	// is the position of result map, starting at 0
	i := 0

	//
	result := make([]map[string]interface{}, TestCreatedSlices)

	//
	for i < TestCreatedSlices {
		which := i % testFixedNumberOfOrderTypes
		otype := TestOrderTypes(which).String()
		t.Logf("%d (of %d) mod %d = %s(%d)", TestZeroIndexQueryMap, TestCreatedSlices, testFixedNumberOfOrderTypes, otype, which)
		if which == 0 {
			t.Logf("QMINDEX(0): %d", TestZeroIndexQueryMap)
			result[i] = TestQueryMap[TestZeroIndexQueryMap]
			result[i]["type"] = otype
			TestZeroIndexQueryMap++
		} else {
			t.Logf("QMINDEX: %d", TestZeroIndexQueryMap-1)
			result[i] = map[string]interface{}{
				"customer_id": TestQueryMap[TestZeroIndexQueryMap-1]["customer_id"],
				"id":          TestQueryMap[TestZeroIndexQueryMap-1]["id"],
				"type":        otype,
				"asset":       "EUR",
			}
		}
		t.Logf("%d of %d (l:%d)", i+1, TestCreatedSlices, len(result[i]))
		ShowSize(result[i])
		//MockRewrite(otype, result, TestQueryMap, TestZeroIndexQueryMap)
		i++
	}

	t.Logf("%v", result)

}

//
func MockRewrite(o string, r []map[string]interface{}, q []map[string]interface{}, i int) {

	for k, v := range r[i] {

		log.Printf("TestZeroIndexQueryMap(%d):OPTYPE(%s):KEY(%s):VALUE(%v)", i, o, k, v)

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

	}

}

// TestForLoopWithOrderTypes
// cmd: go test -v accounting_test.go -run TestForLoopWithOrderTypes
func TestForLoopWithOrderTypes(t *testing.T) {

	t.Run("Results", func(t *testing.T) {

		// is the position of current map, starting at 0
		TestZeroIndexQueryMap := 0
		for TestZeroIndexQueryMap < TestCreatedSlices {

			which := TestZeroIndexQueryMap % testFixedNumberOfOrderTypes
			t.Logf("%d:%s", TestZeroIndexQueryMap, TestOrderTypes(which).String())

			if which == 0 {
				which++
			}

			TestZeroIndexQueryMap++

		}

	})

}

/*
  Test management, V
  add utilities
*/

// is a string array, with the capacity of 5, holding suffixes
// (Bytes, KB, MB, GB, TB) for given calculated values
var suffixes [5]string

// GetSliceHeader inspecting the header values of each slice
func GetSliceHeader(slice *interface{}) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

// GetSize obtains the size of v using reflect and then, for the supported types
// (slices, maps, strings, and structs), it computes the memory required by the content
// stored in them. You would need to add here other types that you need to support.
//
// There are a few details to work out:
// - Private fields are not counted.
// - For structs we are double-counting the basic types.
//
// For number two, filter them out before doing the recursive call when handling structs,
// also check the kinds in the documentation for the reflect package.
func GetSize(v interface{}) int {

	size := int(reflect.TypeOf(v).Size())

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			size += GetSize(s.Index(i).Interface())
		}

	case reflect.Map:

		s := reflect.ValueOf(v)
		keys := s.MapKeys()
		size += int(float64(len(keys)) * 10.79) // approximation from https://golang.org/src/runtime/hashmap.go
		for i := range keys {
			size += GetSize(keys[i].Interface()) + GetSize(s.MapIndex(keys[i]).Interface())
		}

	case reflect.String:
		size += reflect.ValueOf(v).Len()

	case reflect.Struct:
		s := reflect.ValueOf(v)
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).CanInterface() {
				size += GetSize(s.Field(i).Interface())
			}
		}

	}
	return size
}

// ShowSize ...
func ShowSize(result interface{}) {

	suffixes[0] = "Bytes"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	size, _ := strconv.ParseFloat(strconv.Itoa(GetSize(result)), 64)
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	log.Printf(strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix))

}

// Round ...
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
