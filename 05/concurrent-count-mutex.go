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
	count *int,
	mu *sync.Mutex,
) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			num.ProcessingPrimeNumbers(i)
			fmt.Println("Robot", robotNumber, "found prime number:", i)

			mu.Lock()
			*count++
			fmt.Println("total prime numbers found so far:", *count)
			mu.Unlock()
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()
	count := 0
	var mu sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(6)

	go FindPrimeNumbers(1, 100, &wg, 1, &count, &mu)
	go FindPrimeNumbers(101, 200, &wg, 2, &count, &mu)
	go FindPrimeNumbers(201, 300, &wg, 3, &count, &mu)
	go FindPrimeNumbers(301, 400, &wg, 4, &count, &mu)
	go FindPrimeNumbers(401, 500, &wg, 5, &count, &mu)
	go FindPrimeNumbers(501, 600, &wg, 6, &count, &mu)

	wg.Wait()

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}

/*

- SYNCHRONIZATION
  - memory sharing => mutex
  - message passing => channel

*/
