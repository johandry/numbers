package number

import "testing"

func TestPrimeIsPrime(t *testing.T) {
	var tests = []struct {
		n              int
		expectedResult bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{53, true},
		{12345, false},
	}

	for _, test := range tests {
		if actualResult := Prime(test.n).IsPrime(); actualResult != test.expectedResult {
			t.Errorf("Expected %t for number %d, but got %t", test.expectedResult, test.n, actualResult)
		}
	}
}
