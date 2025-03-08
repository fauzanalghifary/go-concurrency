package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"sync"
	"time"
)

func FindPrimeNumbers(lowerBound, upperBound int, wg *sync.WaitGroup, robotNumber int, count *int) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			time.Sleep(100 * time.Millisecond) // Simulate a long-running operation
			fmt.Println("Robot", robotNumber, "found prime number:", i)

			*count++
			fmt.Println("total prime numbers found so far:", *count)
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()
	count := 0

	wg := sync.WaitGroup{}
	wg.Add(8)

	go FindPrimeNumbers(1, 100, &wg, 1, &count)
	go FindPrimeNumbers(101, 200, &wg, 2, &count)
	go FindPrimeNumbers(201, 300, &wg, 3, &count)
	go FindPrimeNumbers(301, 400, &wg, 4, &count)
	go FindPrimeNumbers(401, 500, &wg, 5, &count)
	go FindPrimeNumbers(501, 600, &wg, 6, &count)
	go FindPrimeNumbers(601, 700, &wg, 7, &count)
	go FindPrimeNumbers(701, 800, &wg, 8, &count)

	wg.Wait()

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}

/*
- SYNCHRONIZATION
*/
