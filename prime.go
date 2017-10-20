package number

import (
	"fmt"
)

// Prime represents an integer
type Prime int

// IsPrime returns true if the number is prime
func (p Prime) IsPrime() bool {
	if (int(p) == 1) || (int(p) == 2) {
		return true
	}
	return p.isPrimeSieveOfEratosthenesSec()
}

// IsPrime resturns true if the given integer is prime
func IsPrime(n int) bool {
	return Prime(n).IsPrime()
}

// Sieve is an set of booleans meaning that if element at a given index (i) is
// true the number (2*i + 3) is prime
type Sieve []bool

// newSieve creates a Sieve for odd numbers (even numbers are not prime) assuming
// all of them are prime by default, so the length of the sieve is half the number
func newSieve(n int) Sieve {
	sieve := make(Sieve, n/2)
	for i := 0; i < n/2; i++ {
		sieve[i] = true
	}
	return sieve
}

// valueAt returns the number is the sieve that should be in the given index
func valueAt(index int) int {
	return index + index + 3
}

func (s Sieve) String() string {
	var list []int
	for i, h := range s {
		if h {
			list = append(list, valueAt(i))
		}
	}
	return fmt.Sprint(list)
}

// mark marks all the numbers in a sive from init to end for a given factor as
// non-prime (mark it as false every factor number)
func (s *Sieve) mark(init, end, factor int) {
	for i := init; i < end && i < len(*s); i += factor {
		(*s)[i] = false
		fmt.Println(valueAt(i), "is not a prime")
	}
}

func (s Sieve) primes() []Prime {
	primes := make([]Prime, 0)
	for i, h := range s {
		if h {
			primes = append(primes, Prime(valueAt(i)))
		}
	}
	return primes
}

// LowerPrimes01 returns all the prime numbers lower than prime number p
// Non-optimal option #1:
// starting in the next factor and jumpping every factor times until get to
// the last element. f(53) uses 25 loops
func (p Prime) LowerPrimes01() []Prime {
	sieve := newSieve(int(p))
	var factor int
	for i := 0; i < len(sieve); i++ {
		factor = valueAt(i)
		sieve.mark(i+factor, len(sieve), factor)
		fmt.Printf("i:%d, value at #i: %d, factor: %d, Sieve: %s\n", i, valueAt(i), factor, sieve)
	}

	return sieve.primes()
}

func indexSqrt(i int) int {
	// 2i^2 + 6i + 3 == 2i * (i + 3) + 3
	return 2*i*(i+3) + 3
}

// LowerPrimes02 returns all the prime numbers lower than prime number p
// Non-optimal option #2:
// Starting in the square of first prime number, jumpping every 2*i times for
// square_root(len(p)) times
func (p Prime) LowerPrimes02() []Prime {
	sieve := newSieve(int(p))
	iSqrt := 3
	for i := 0; iSqrt < int(p); i++ {
		if sieve[i] {
			sieve.mark(i+iSqrt, i+int(p), i+i+3)
		}
		fmt.Printf("i:%d, value at #i: %d, iSqrt: %d, Sieve: %s\n", i, valueAt(i), iSqrt, sieve)
		iSqrt = indexSqrt(i)
	}

	return sieve.primes()
}

// LowerPrimes returns all the prime numbers lower than prime number p
// Optimal option #3:
// Same as option #2 but iSqrt = factor(i) + factor(i+1) and factor += 2
func (p Prime) LowerPrimes() []Prime {
	sieve := newSieve(int(p))

	last := int(p)
	iSqrt := 3
	factor := 3
	for i := 0; iSqrt < int(p); i++ {
		// iSqrt = 2i^2 + 6i + 3
		// factor = 2i + 3
		if sieve[i] {
			sieve.mark(i+iSqrt, last, factor)
		}
		fmt.Printf("i:%d, value at #i: %d, iSqrt: %d, factor: %d, Sieve: %s\n", i, valueAt(i), iSqrt, factor, sieve)
		iSqrt += factor
		factor += 2
		iSqrt += factor
	}

	return sieve.primes()
}

// LowerPrimes returns all the prime numbers lower than integer n
func LowerPrimes(n int) []Prime {
	return Prime(n).LowerPrimes()
}

// isPrimeSieveOfEratosthenesSec returns true if the number p is Prime
// It gets the list of primes lower than p *Secuencially* and try to find it there
func (p Prime) isPrimeSieveOfEratosthenesSec() bool {
	lowerPs := p.LowerPrimes()
	fmt.Printf("Primes lower than %d: %v\n", int(p), lowerPs)
	for i := len(lowerPs) - 1; int(lowerPs[i]) >= int(p); i-- {
		if int(lowerPs[i]) == int(p) {
			return true
		}
	}
	return false
}
