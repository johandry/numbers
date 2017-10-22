package number

import (
	"fmt"
	"strconv"
	"strings"
)

// Sieve is an set of booleans meaning that if element at a given index (i) is
// true the number (2*i + 3) is prime
type Sieve []bool

// newSieve creates a Sieve for odd numbers (even numbers are not prime) assuming
// all of them are prime by default, so the length of the sieve is half the number
func newSieve(n int) Sieve {
	if n <= 2 {
		return Sieve{}
	}
	l := (n - 1) / 2
	sieve := make(Sieve, l)
	for i := range sieve {
		sieve[i] = true
	}
	return sieve
}

// Debug enable the debug mode when set to true
var Debug bool

const sieveMark = "\x1b[38;5;8m"
const sieveMarkReset = "\x1b[0m"

func debug(format string, args ...interface{}) {
	if Debug {
		format = strings.TrimRight(format, "\n")
		format = strings.Replace(format, "\n", "\n        ", -1)
		fmt.Printf("\x1b[35m[DEBUG]\x1b[0m "+format+"\n", args...)
	}
}

// valueAt returns the number is the sieve that should be in the given index
func valueAt(index int) int {
	return index + index + 3
}

func indexOf(n int) int {
	return (n - 3) / 2
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

// StringDebug converts the sieve to a string but showing the marked numbers
func (s Sieve) StringDebug() string {
	var list []string
	for i, h := range s {
		if h {
			list = append(list, strconv.Itoa(valueAt(i)))
		} else {
			list = append(list, sieveMark+strconv.Itoa(valueAt(i))+sieveMarkReset)
		}
	}
	return "[" + strings.Join(list, " ") + "]"
}

// mark marks all the numbers in a sive from init to end for a given factor as
// non-prime (mark it as false every factor number)
func (s *Sieve) mark(init, end, factor int) {
	// end := valueAt(len(*s))
	(*s)[init] = false
	// if valueAt(init)%factor != 0 {
	// 	debug("%d was marked but is not multiple of factor %d", valueAt(init), factor)
	// }
	for end-init > factor {
		init += factor
		(*s)[init] = false
		// if valueAt(init)%factor != 0 {
		// 	debug("%d was marked but is not multiple of factor %d", valueAt(init), factor)
		// }
	}
}

// Sift sifts the sieve to crossout all the non-prime numbers
// Optimal option #3:
// Same as option #2 but iSqrt = factor(i) + factor(i+1) and factor += 2
func (s *Sieve) Sift() {
	if len(*s) == 0 {
		return
	}
	n := len(*s)
	// debug("N: %d", n)
	iSqrt := 3
	factor := 3
	for i := 0; iSqrt < n; i++ {
		// iSqrt = 2i^2 + 6i + 3
		// factor = 2i + 3
		// debug("i:%d, value at #i: %d, iSqrt: %d, factor: %d\nSieve before mark: %s", i, valueAt(i), iSqrt, factor, s.StringDebug())
		debug("i:%d, value at #i: %d, iSqrt: %d, factor: %d", i, valueAt(i), iSqrt, factor)
		if (*s)[i] {
			s.mark(iSqrt, n, factor)
		}
		debug("Sieve after mark:  %s", s.StringDebug())
		// factor = i + i + 3
		// iSqrt = 2*i*(i+3) + 3
		iSqrt += factor
		factor += 2
		iSqrt += factor
	}
}

func indexSqrt(i int) int {
	// 2i^2 + 6i + 3 == 2i * (i + 3) + 3
	return 2*i*(i+3) + 3
}

// Sift01 sifts the sieve to crossout all the non-prime numbers using other algorithm
// Non-optimal option #1:
// starting in the next factor and jumpping every factor times until get to
// the last element. f(53) uses 25 loops
func (s *Sieve) Sift01() {
	factor := 3
	for i := 0; factor*factor < valueAt(len(*s)); {
		if (*s)[i] {
			s.mark(i+factor, len(*s), factor)
			debug("i:%d, value at #i: %d, factor: %d, Sieve: %s", i, valueAt(i), factor, s.StringDebug())
		}
		i++
		factor = valueAt(i)
	}
}

// Sift02 sifts the sieve to crossout all the non-prime numbers using other algorithm
// Non-optimal option #2:
// Starting in the square of first prime number, jumpping every 2*i times for
// square_root(len(p)) times
func (s *Sieve) Sift02() {
	n := len(*s)
	iSqrt := 3
	for i := 0; iSqrt < n; {
		if (*s)[i] {
			s.mark(iSqrt, n, i+i+3)
		}
		debug("i:%d, value at #i: %d, iSqrt: %d, Sieve: %s", i, valueAt(i), iSqrt, s.StringDebug())
		i++
		iSqrt = indexSqrt(i)
	}
}

// Primes return all the prime numbers in a sieve
func (s Sieve) Primes() []Prime {
	primes := []Prime{Prime(2)}
	for i, h := range s {
		if h {
			primes = append(primes, Prime(valueAt(i)))
		}
	}
	return primes
}
