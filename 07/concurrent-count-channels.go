package main

import (
	num "fauzanalghifary/go-concurrency"
	"fmt"
	"sync"
	"time"
)

func FindPrimeNumbers(
	lowerBound, upperBound int,
	wg *sync.WaitGroup,
	robotNumber int,
	results chan<- int,
) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			num.ProcessingPrimeNumbers(i)
			fmt.Println("Robot", robotNumber, "found prime number:", i)

			results <- i
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()

	results := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(6)

	go FindPrimeNumbers(1, 100, &wg, 1, results)
	go FindPrimeNumbers(101, 200, &wg, 2, results)
	go FindPrimeNumbers(201, 300, &wg, 3, results)
	go FindPrimeNumbers(301, 400, &wg, 4, results)
	go FindPrimeNumbers(401, 500, &wg, 5, results)
	go FindPrimeNumbers(501, 600, &wg, 6, results)

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for prime := range results {
		count++
		fmt.Println("count increased from", count-1, "to", count, "with prime:", prime)
	}

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}

/*

- SYNCHRONIZATION
  - memory sharing => mutex
  - message passing => channel

*/
