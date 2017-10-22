package number

import "testing"

var sieveNewDataTests = []struct {
	n              int
	expectedResult Sieve
}{
	{0, nil},
	{1, nil},
	{2, nil},
	{3, Sieve{true}},
	{4, Sieve{true}},
	{5, Sieve{true, true}},
	{6, Sieve{true, true}},
	{7, Sieve{true, true, true}},
	{8, Sieve{true, true, true}},
	{9, Sieve{true, true, true, true}},
}

func equalSieves(s1, s2 Sieve) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestNewSieve(t *testing.T) {
	for _, test := range sieveNewDataTests {
		if actualResult := newSieve(test.n); !equalSieves(actualResult, test.expectedResult) {
			t.Errorf("Expected %v for %d-Sieve, but got %v", test.expectedResult, test.n, actualResult)
		}
	}
}

var sieveStringDataTests = []struct {
	n              int
	expectedResult string
}{
	{0, "[]"},
	{1, "[]"},
	{2, "[]"},
	{3, "[3]"},
	{4, "[3]"},
	{5, "[3 5]"},
	{6, "[3 5]"},
	{7, "[3 5 7]"},
	{8, "[3 5 7]"},
	{9, "[3 5 7]"},
	{11, "[3 5 7 11]"},
	{12, "[3 5 7 11]"},
	{13, "[3 5 7 11 13]"},
	{14, "[3 5 7 11 13]"},
}

func TestStringSieve(t *testing.T) {
	var s Sieve
	for _, test := range sieveStringDataTests {
		s = newSieve(test.n)
		s.Sift()
		if actualResult := s.String(); actualResult != test.expectedResult {
			t.Errorf("Expected %s for sifted %d-Sieve, but got %s", test.expectedResult, test.n, actualResult)
		}
	}
}

func TestStringSieve01(t *testing.T) {
	var s Sieve
	for _, test := range sieveStringDataTests {
		s = newSieve(test.n)
		s.Sift01()
		if actualResult := s.String(); actualResult != test.expectedResult {
			t.Errorf("Expected %s for sifted %d-Sieve, but got %s", test.expectedResult, test.n, actualResult)
		}
	}
}

func TestStringSieve02(t *testing.T) {
	var s Sieve
	for _, test := range sieveStringDataTests {
		s = newSieve(test.n)
		s.Sift02()
		if actualResult := s.String(); actualResult != test.expectedResult {
			t.Errorf("Expected %s for sifted %d-Sieve, but got %s", test.expectedResult, test.n, actualResult)
		}
	}
}
