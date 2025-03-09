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
			num.ProcessingPrimeNumbers(i)
			fmt.Println("Robot", robotNumber, "found prime number:", i)
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(6)

	// spawn 6 goroutines
	go FindPrimeNumbers(1, 100, &wg, 1)
	go FindPrimeNumbers(101, 200, &wg, 2)
	go FindPrimeNumbers(201, 300, &wg, 3)
	go FindPrimeNumbers(301, 400, &wg, 4)
	go FindPrimeNumbers(401, 500, &wg, 5)
	go FindPrimeNumbers(501, 600, &wg, 6)

	wg.Wait()
	fmt.Println("finished in:", time.Since(start))
}

/*

- Concurrent program
- Concurrency vs Parallelism
  - Concurrency is about PLANNING how to do many tasks at the same time;
  - Parallelism is about PERFORMING many tasks at the same time.
  - Concurrency ENABLES parallelism.

- The requirement of parallel execution:
  - Hardware support
  - Task independence
	~ Task decomposition
	~ Data decomposition

*/
