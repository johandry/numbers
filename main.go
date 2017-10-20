package main

import (
	"fmt"

	"github.com/johandry/numbers/prime"
)

func main() {
	p := prime.Prime(53)
	if p.IsPrime() {
		fmt.Printf("%d is Prime\n", p)
	} else {
		fmt.Printf("%d is Composite\n", p)
	}
}
