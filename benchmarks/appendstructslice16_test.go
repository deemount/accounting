package benchmarks

import (
	"testing"

	"github.com/google/uuid"
)

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
