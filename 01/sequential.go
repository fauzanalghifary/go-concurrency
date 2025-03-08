package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"time"
)

func FindPrimeNumbers(lowerBound, upperBound int) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			time.Sleep(100 * time.Millisecond) // Simulate a long-running operation
			fmt.Println("Found prime number: ", i)
		}
	}
}

func main() {
	start := time.Now()

	FindPrimeNumbers(1, 800)

	fmt.Println("finished in:", time.Since(start))
}
