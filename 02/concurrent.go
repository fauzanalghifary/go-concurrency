package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"sync"
	"time"
)

func FindPrimeNumbers(lowerBound, upperBound int, wg *sync.WaitGroup, robotNumber int) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			time.Sleep(100 * time.Millisecond) // Simulate a long-running operation
			fmt.Println("Robot", robotNumber, "found prime number:", i)
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(8)

	go FindPrimeNumbers(1, 100, &wg, 1)
	go FindPrimeNumbers(101, 200, &wg, 2)
	go FindPrimeNumbers(201, 300, &wg, 3)
	go FindPrimeNumbers(301, 400, &wg, 4)
	go FindPrimeNumbers(401, 500, &wg, 5)
	go FindPrimeNumbers(501, 600, &wg, 6)
	go FindPrimeNumbers(601, 700, &wg, 7)
	go FindPrimeNumbers(701, 800, &wg, 8)

	wg.Wait()

	fmt.Println("finished in:", time.Since(start))
}
