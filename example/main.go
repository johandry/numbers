package main

import (
	"fmt"

	"github.com/johandry/number"
)

func main() {
	// n := 52
	// if number.IsPrime(n) {
	// 	fmt.Printf("%d is Prime\n", n)
	// } else {
	// 	fmt.Printf("%d is Composite\n", n)
	// }

	p := number.Prime(53)
	if p.IsPrime() {
		fmt.Printf("%d is Prime\n", p)
	} else {
		fmt.Printf("%d is Composite\n", p)
	}
}
