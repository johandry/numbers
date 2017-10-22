package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/johandry/number"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, "Enter at least a number\n")
		usage(1)
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage(0)
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-d" {
			number.Debug = !number.Debug
			continue
		}
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s is not a valid number. %s\n", os.Args[i], err)
			continue
		}
		if number.IsPrime(n) {
			fmt.Printf("%d is Prime\n", n)
		} else {
			fmt.Printf("%d is Composite\n", n)
		}
	}
}

func usage(exitCode int) {
	fmt.Println("Usage: Enter a list of integers to print if they are prime or composite.")
	fmt.Println("       If used '-d' instead of a number of the list, it will toggle the debug mode")
	fmt.Println("Example: -d 1 2 3 4 5 6 7 -d 8 9 10 11 12 13 14")
	os.Exit(exitCode)
}
