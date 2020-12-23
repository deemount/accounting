package tests

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

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

// TestOrderTypesValues
// cmd: go test -v accounting_test.go -run TestOrderTypesValues
func TestOrderTypesValues(t *testing.T) {

	// show integer values of types
	t.Run("Results", func(t *testing.T) {
		t.Logf("Withdrawal: %v\n", TestOrderTypes(0))
		t.Logf("Buy: %v\n", TestOrderTypes(1))
		t.Logf("Spread: %v\n", TestOrderTypes(2))
		t.Logf("Fee: %v\n", TestOrderTypes(3))
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

	// reflect value of order types
	// first use case fails as given example for better understanding
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

		ShowSize(result)
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
			ShowSize(result[i])
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
