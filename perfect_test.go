package number

import "testing"

var isPerfectDataTest = []struct {
	n              int
	expectedResult bool
}{
	{1, false},
	{2, false},
	{3, false},
	{5, false},
	{6, true},
	{8, false},
	{28, true},
	{120, false},
	{496, true},
	{2016, false},
	{8128, true},
	{8130, false},
	{33550336, true},
	{33550338, false},
	// {8589869054, false},
	{8589869056, true},
	// {137438691328, true},
	// {2305843008139952128, true},
}

func TestIsPerfect(t *testing.T) {
	for _, test := range isPerfectDataTest {
		if actualResult := IsPerfect(test.n); actualResult != test.expectedResult {
			t.Errorf("Expected %t for number %d, but got %t", test.expectedResult, test.n, actualResult)
		}
	}
}
