package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"time"
)

func FindPrimeNumbers(lowerBound, upperBound int, count *int) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			time.Sleep(100 * time.Millisecond) // Simulate a long-running operation
			fmt.Println("Found prime number: ", i)

			*count++
			fmt.Println("total prime numbers found so far:", *count)
		}
	}
}

func main() {
	start := time.Now()
	count := 0

	FindPrimeNumbers(1, 800, &count)

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}
