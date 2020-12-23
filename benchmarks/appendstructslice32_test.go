package benchmarks

import (
	"testing"

	"github.com/google/uuid"
)

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
