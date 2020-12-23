package tests

import "github.com/google/uuid"

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
