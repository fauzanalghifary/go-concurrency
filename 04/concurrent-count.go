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
			num.ProcessingPrimeNumbers(i)
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
	wg.Add(6)

	go FindPrimeNumbers(1, 100, &wg, 1, &count)
	go FindPrimeNumbers(101, 200, &wg, 2, &count)
	go FindPrimeNumbers(201, 300, &wg, 3, &count)
	go FindPrimeNumbers(301, 400, &wg, 4, &count)
	go FindPrimeNumbers(401, 500, &wg, 5, &count)
	go FindPrimeNumbers(501, 600, &wg, 6, &count)

	wg.Wait()

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}

/*

- RACE CONDITION
  - two or more goroutines access the same variable
  - at least one of them is doing write operation
- SYNCHRONIZATION
  - memory sharing
  - message passing

*/
