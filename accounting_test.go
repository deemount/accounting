package accounting

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"testing"
	"unsafe"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestOrderTypes int64

const (
	TestWithdrawal TestOrderTypes = iota
	TestBuy
	TestSpread
	TestFee
)

// String return name of typed constants by given integer (0,1,2,3,4)
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

/**/

// TestExchangeOrder ...
type TestExchangeOrder struct {
	ID         uuid.UUID `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Type       string    `json:"type"`
	Asset      string    `json:"asset"`
}

// Private
// otype: is string value
// which: is integer value calculated with modulo and represents\nthe value of a typed constant

var (

	//
	testQueryMap []map[string]interface{}

	// is a string array, with the capacity of 5, holding suffixes (Bytes, KB, MB, GB, TB) for given calculated values
	suffixes [5]string

	// is number of wanted testFixedNumberOfAdditionalSlices for each map in map slice
	testFixedNumberOfAdditionalSlices = 4

	// filled single struct for testing
	testDataStruct = TestExchangeOrder{
		ID:         uuid.New(),
		CustomerID: 1,
		Type:       "buy",
		Asset:      "EUR",
	}

	// filled struct slice for testing
	testDataStructSlice = []TestExchangeOrder{
		{
			ID:         uuid.New(),
			CustomerID: 1,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 2,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 3,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 4,
			Type:       "buy",
			Asset:      "EUR",
		},
	}

	//
	testNumberOfSlicesInTestStruct = len(testDataStructSlice)

	// is number of all maps by multiplication of testNumberOfSlicesInTestStruct and testFixedNumberOfAdditionalSlices
	testCreatedSlices = testNumberOfSlicesInTestStruct * testFixedNumberOfAdditionalSlices
)

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
	t.Run("PointerTo", func(t *testing.T) {

		rt3 := reflect.TypeOf(TestOrderTypes(0))
		rt4 := reflect.TypeOf(new(TestOrderTypes))

		ok := assert.Equal(t, reflect.PtrTo(rt3), rt4)
		if ok {
			t.Log("is equal, because both have same type of pointer")
		}
	})
	t.Log("\n\n")

	//
	t.Run("ValueOf", func(t *testing.T) {

		rt5 := reflect.ValueOf(TestOrderTypes(1))
		rt6 := reflect.ValueOf(int64(1))

		ok1 := assert.Equal(t, rt5, rt6)
		if !ok1 {
			t.Log("reflect.ValueOf(TestOrderTypes(1)) and reflect.ValueOf(int64(1)) is not equal")
		}

		ok2 := assert.Equal(t, reflect.Kind(TestOrderTypes(1)), reflect.Kind(int64(1)))
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

		bytes, _ := json.Marshal(&testDataStruct)
		json.Unmarshal(bytes, &q1)

		t.Logf("Single: %+v", q1)

	})
	t.Log("\n\n")

	//
	t.Run("Many", func(t *testing.T) {

		q2 := make([]map[string]interface{}, testNumberOfSlicesInTestStruct)

		bytes, _ := json.Marshal(&testDataStructSlice)
		json.Unmarshal(bytes, &q2)

		t.Logf("Many: %v", q2)

	})

}

// TestMakeMaps
// cmd: go test -v accounting_test.go -run TestMakeMaps
func TestMakeMaps(t *testing.T) {

	t.Run("Many", func(t *testing.T) {

		emptyMaps := make([]map[string]interface{}, testNumberOfSlicesInTestStruct)
		t.Logf("%v", emptyMaps)

	})

}

// TestAppendStructSlice
// cmd: go test -v accounting_test.go -run TestAppendStructSlice
func TestAppendStructSlice(t *testing.T) {

	t.Run("Show", func(t *testing.T) {

		// is the position of current map, starting at 0
		testZeroIntegerIndex := 0
		result := make([]TestExchangeOrder, testCreatedSlices)
		for i := 0; i < testCreatedSlices; i++ {
			if i%4 == 0 {
				result[i] = testDataStructSlice[testZeroIntegerIndex]
				testZeroIntegerIndex++
			}
		}

		showSize(result)
		t.Logf("%d:%d", len(result), cap(result))
		b, _ := json.MarshalIndent(result, " ", "   ")
		t.Logf("%v", string(b))

	})

}

// TestAppendMapSlice
// cmd: go test -v accounting_test.go -run TestAppendMapSlice
func TestAppendMapSlice(t *testing.T) {

	t.Run("Show", func(t *testing.T) {

		b, _ := json.Marshal(&testDataStructSlice)
		json.Unmarshal(b, &testQueryMap)

		// is the position of current map, starting at 0
		testZeroIntegerIndex := 0
		result := make([]map[string]interface{}, testCreatedSlices)
		for i := 0; i < testCreatedSlices; i++ {
			if i%testFixedNumberOfAdditionalSlices == 0 {
				result[i] = testQueryMap[testZeroIntegerIndex]
				testZeroIntegerIndex++
			}
		}

		showSize(result)
		t.Logf("%d:%d", len(result), cap(result))
		b1, _ := json.MarshalIndent(result, " ", "   ")
		t.Logf("%v", string(b1))

	})

}

// TestAppendRewriteMapSlice
// cmd: go test -v accounting_test.go -run TestAppendRewriteMapSlice
func TestAppendRewriteMapSlice(t *testing.T) {

	t.Run("Slice", func(t *testing.T) {

		b, _ := json.Marshal(&testDataStructSlice)
		json.Unmarshal(b, &testQueryMap)

		// is the position of current map, starting at 0
		testZeroIntegerIndex := 0
		result := make([]map[string]interface{}, testCreatedSlices)
		for i := 0; i < testCreatedSlices; i++ {
			which := i % testFixedNumberOfAdditionalSlices
			if which == 0 {
				result[i] = testQueryMap[testZeroIntegerIndex]
				result[i]["type"] = TestOrderTypes(which).String()
				testZeroIntegerIndex++
			} else {
				result[i] = map[string]interface{}{
					"customer_id": testQueryMap[testZeroIntegerIndex-1]["customer_id"],
					"id":          testQueryMap[testZeroIntegerIndex-1]["id"],
					"type":        TestOrderTypes(which).String(),
					"asset":       "EUR",
				}
			}
			t.Logf("%d of %d (l:%d)", i, testCreatedSlices, len(result[i]))
			showSize(result[i])
		}

	})

}

func TestAll(t *testing.T) {

	// is the position of current map, starting at 0
	testZeroIntegerIndex := 0
	for testZeroIntegerIndex < testCreatedSlices {
		which := testZeroIntegerIndex % testFixedNumberOfAdditionalSlices
		//otype := TestOrderTypes(which).String()
		//t.Logf("%d (of %d) mod %d = %s(%d)", testZeroIntegerIndex, testCreatedSlices, testFixedNumberOfAdditionalSlices, otype, which)
		//MockRewrite(otype, result, testQueryMap, testZeroIntegerIndex)
		if which == 0 {
			which++
		}
		testZeroIntegerIndex++
	}

}

//
func MockRewrite(o string, r []map[string]interface{}, q []map[string]interface{}, i int) {

	for k, v := range r[i] {

		log.Printf("testZeroIntegerIndex(%d):OPTYPE(%s):KEY(%s):VALUE(%v)", i, o, k, v)

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
		testZeroIntegerIndex := 0
		for testZeroIntegerIndex < testCreatedSlices {

			which := testZeroIntegerIndex % testFixedNumberOfAdditionalSlices
			t.Logf("%d:%s", testZeroIntegerIndex, TestOrderTypes(which).String())

			if which == 0 {
				which++
			}

			testZeroIntegerIndex++

		}

	})

}

/*

Run all benchmark tests without the other ttesting functions with this cmd:
go test -run=^$ -bench=.

*/

// BenchmarkAppendStructSlice1
func BenchmarkAppendStructSlice1(b *testing.B) {
	var testDataStructSlice1 = []TestExchangeOrder{
		{
			ID:         uuid.New(),
			CustomerID: 1,
			Type:       "buy",
			Asset:      "EUR",
		},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testZeroIntegerIndex1 := 0
		result1 := make([]TestExchangeOrder, 4)
		for i := 0; i < 4; i++ {
			if i%4 == 0 {
				result1[i] = testDataStructSlice1[testZeroIntegerIndex1]
				testZeroIntegerIndex1++
			}
		}
	}
}

// BenchmarkAppendStructSlice16
func BenchmarkAppendStructSlice16(b *testing.B) {
	var testDataStructSlice1 = []TestExchangeOrder{
		{
			ID:         uuid.New(),
			CustomerID: 1,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 2,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 3,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 4,
			Type:       "buy",
			Asset:      "EUR",
		},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testZeroIntegerIndex1 := 0
		result1 := make([]TestExchangeOrder, 16)
		for i := 0; i < 16; i++ {
			if i%4 == 0 {
				result1[i] = testDataStructSlice1[testZeroIntegerIndex1]
				testZeroIntegerIndex1++
			}
		}
	}
}

// BenchmarkAppendStructSlice16
func BenchmarkAppendStructSlice32(b *testing.B) {
	var testDataStructSlice1 = []TestExchangeOrder{
		{
			ID:         uuid.New(),
			CustomerID: 1,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 2,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 3,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 4,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 5,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 6,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 7,
			Type:       "buy",
			Asset:      "EUR",
		},
		{
			ID:         uuid.New(),
			CustomerID: 8,
			Type:       "buy",
			Asset:      "EUR",
		},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testZeroIntegerIndex1 := 0
		result1 := make([]TestExchangeOrder, 32)
		for i := 0; i < 32; i++ {
			if i%4 == 0 {
				result1[i] = testDataStructSlice1[testZeroIntegerIndex1]
				testZeroIntegerIndex1++
			}
		}
	}
}

// getSliceHeader
// inspecting the header values of each slice
func getMapSliceHeader(slice *[]map[string]interface{}) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

// getSliceHeader
// inspecting the header values of each slice
func getSliceHeader(slice *map[string]interface{}) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
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

func showSize(result interface{}) {

	suffixes[0] = "Bytes"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	size, _ := strconv.ParseFloat(strconv.Itoa(getSize(result)), 64)
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	log.Printf(strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix))

}

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
