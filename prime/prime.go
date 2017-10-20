package prime

import "fmt"

// Prime represents an integer
type Prime int

// IsPrime returns true if the number is prime
func (p Prime) IsPrime() bool {
	return p.isPrimeSieveOfEratosthenesSec()
}

func newSieve(l int) []int {
	sieve := make([]int, l/2)
	for i := 0; i < l/2; i++ {
		sieve[i] = 2*i + 3
	}
	return sieve
}

func (p Prime) isPrimeSieveOfEratosthenesSec() bool {
	sieve := newSieve(int(p))
	fmt.Println(sieve)
	return false
}
