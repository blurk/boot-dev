package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			println(i)
			continue
		}

		if i%2 == 0 {
			continue
		}

		isPrime := true

		for j := 3; j < int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			println(i)
		}

	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
