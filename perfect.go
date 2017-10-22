package number

import "math"

// Perfect represents a possible perfect number
type Perfect int

// IsPerfect returns true is the number is perfect
func (p Perfect) IsPerfect() bool {
	// This may not be true for numbers larger than 10^1500 but the larger uint64
	// is 20^20, so given p it won't be perfect if it's odd
	if int(p)%2 != 0 {
		return false
	}
	return p.isPerfectEuclidSec()
}

// IsPerfect returns true is the number is perfect
func IsPerfect(n int) bool {
	return Perfect(n).IsPerfect()
}

// func (p Perfect) isPerfect01() bool {
// 	var sum01, sum02 float64
// 	var i int
// 	for i = 1; i < int(p)-2; i++ {
// 		sum01 += float64(i) * (float64(p) / float64(i))
// 		sum02 += float64(i) * ((float64(p) - 1) / float64(i))
// 	}
// 	i++
// 	sum02 += float64(i) * ((float64(p) - 1) / float64(i))
// 	debug("%f - %f", sum01, sum02)
// 	return sum01 == sum02
// }

func euclids(n int) int {
	// 2^(n−1) * (2^n − 1)
	return int(math.Pow(2, float64(n-1)) * (math.Pow(2, float64(n)) - 1))
}
func (p Perfect) isPerfectEuclidSec() bool {
	for i := 0; i < int(p); i++ {
		if int(p) == euclids(i) {
			debug("2^%d-1 * (2^%d - 1) = %d ... is %d prime?", i, i, euclids(i), int(math.Pow(2, float64(i))-1))
			return IsPrime(int(math.Pow(2, float64(i)) - 1))
		}
	}
	return false
}
