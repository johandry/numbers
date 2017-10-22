package number

// Prime represents an integer
type Prime int

// IsPrime returns true if the number is prime
func (p Prime) IsPrime() bool {
	return p.isPrimeSieveOfEratosthenesSec()
}

// IsPrime resturns true if the given integer is prime
func IsPrime(n int) bool {
	return Prime(n).IsPrime()
}

// LowerPrimes returns all the prime numbers lower than prime number p
func (p Prime) LowerPrimes() []Prime {
	if int(p) <= 1 {
		return nil
	}
	if int(p) == 2 {
		return []Prime{Prime(2)}
	}
	sieve := newSieve(int(p))
	sieve.Sift()

	return sieve.Primes()
}

// LowerPrimes returns all the prime numbers lower than integer n
func LowerPrimes(n int) []Prime {
	return Prime(n).LowerPrimes()
}

// isPrimeSieveOfEratosthenesSec returns true if the number p is Prime
// It gets the list of primes lower than p *Secuencially* and try to find it there
func (p Prime) isPrimeSieveOfEratosthenesSec() bool {
	if int(p) <= 1 {
		return false
	}
	if int(p) == 2 {
		return true
	}
	if int(p)%2 == 0 {
		return false
	}
	sieve := newSieve(int(p))
	sieve.Sift()

	return sieve[indexOf(int(p))]
}

// Divisors returns a list of prime numbers divisors of the given prime or integer
func (p Prime) Divisors() []Prime {
	lowerPrimes := p.LowerPrimes()
	divisors := []Prime{}
	for _, prime := range lowerPrimes {
		// To not include the number itself as a divisor, replace this line for the next one and modify the tests
		// if int(prime) != 1 && int(prime) != int(p) && int(p)%int(prime) == 0 {
		if int(prime) != 1 && int(p)%int(prime) == 0 {
			divisors = append(divisors, prime)
		}
	}
	return divisors
}

// Divisors returns a list of prime numbers divisors of the given integer
func Divisors(n int) []Prime {
	return Prime(n).Divisors()
}
