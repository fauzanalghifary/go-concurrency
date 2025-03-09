package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"time"
)

func FindPrimeNumbers(lowerBound, upperBound int) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			num.ProcessingPrimeNumbers(i) // Simulate a long-running operation
			fmt.Println("Found prime number: ", i)
		}
	}
}

func main() {
	start := time.Now()

	FindPrimeNumbers(1, 600)

	fmt.Println("finished in:", time.Since(start))
}

/*

- Sequential program
- Executed by the "main goroutine"

*/
