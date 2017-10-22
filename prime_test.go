package number

import "testing"

var isPrimeDataTests = []struct {
	n              int
	expectedResult bool
}{
	{-1, false},
	{0, false},
	{1, false},
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

func TestIsPrime(t *testing.T) {
	for _, test := range isPrimeDataTests {
		if actualResult := IsPrime(test.n); actualResult != test.expectedResult {
			t.Errorf("Expected %t for number %d, but got %t", test.expectedResult, test.n, actualResult)
		}
	}
}

var listOfPrimesDataTests = []struct {
	n              int
	expectedResult []Prime
}{
	{0, nil},
	{1, nil},
	{2, []Prime{Prime(2)}},
	{3, []Prime{Prime(2), Prime(3)}},
	{4, []Prime{Prime(2), Prime(3)}},
	{5, []Prime{Prime(2), Prime(3), Prime(5)}},
	{6, []Prime{Prime(2), Prime(3), Prime(5)}},
	{7, []Prime{Prime(2), Prime(3), Prime(5), Prime(7)}},
	{8, []Prime{Prime(2), Prime(3), Prime(5), Prime(7)}},
	{9, []Prime{Prime(2), Prime(3), Prime(5), Prime(7)}},
	{11, []Prime{Prime(2), Prime(3), Prime(5), Prime(7), Prime(11)}},
	{12, []Prime{Prime(2), Prime(3), Prime(5), Prime(7), Prime(11)}},
	{13, []Prime{Prime(2), Prime(3), Prime(5), Prime(7), Prime(11), Prime(13)}},
	{15, []Prime{Prime(2), Prime(3), Prime(5), Prime(7), Prime(11), Prime(13)}},
	{53, []Prime{Prime(2), Prime(3), Prime(5), Prime(7), Prime(11), Prime(13),
		Prime(17), Prime(19), Prime(23), Prime(29), Prime(31), Prime(37), Prime(41), Prime(43), Prime(47), Prime(53)}},
}

func equalPrimes(pl1, pl2 []Prime) bool {
	if pl1 == nil && pl2 == nil {
		return true
	}
	if pl1 == nil || pl2 == nil {
		return false
	}
	if len(pl1) != len(pl2) {
		return false
	}
	for i := range pl1 {
		if pl1[i] != pl2[i] {
			return false
		}
	}
	return true
}

func TestLowerPrimes(t *testing.T) {
	for _, test := range listOfPrimesDataTests {
		if actualResult := LowerPrimes(test.n); !equalPrimes(actualResult, test.expectedResult) {
			t.Errorf("Expected %v for number %d, but got %v", test.expectedResult, test.n, actualResult)
		}
	}
}

var divisorsDataTest = []struct {
	n              int
	expectedResult []Prime
}{
	{0, []Prime{}},
	{1, []Prime{}},
	{2, []Prime{Prime(2)}},
	{3, []Prime{Prime(3)}},
	{4, []Prime{Prime(2)}},
	{5, []Prime{Prime(5)}},
	{6, []Prime{Prime(2), Prime(3)}},
	{7, []Prime{Prime(7)}},
	{8, []Prime{Prime(2)}},
	{9, []Prime{Prime(3)}},
	{11, []Prime{Prime(11)}},
	{12, []Prime{Prime(2), Prime(3)}},
	{13, []Prime{Prime(13)}},
	{15, []Prime{Prime(3), Prime(5)}},
	{53, []Prime{Prime(53)}},
}

func TestDivisors(t *testing.T) {
	for _, test := range divisorsDataTest {
		if actualResult := Divisors(test.n); !equalPrimes(actualResult, test.expectedResult) {
			t.Errorf("Expected %v for number %d, but got %v", test.expectedResult, test.n, actualResult)
		}
	}
}
