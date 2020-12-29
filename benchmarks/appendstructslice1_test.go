package benchmarks

import (
	"testing"

	"github.com/google/uuid"
)

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
