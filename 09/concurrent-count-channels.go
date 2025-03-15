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
	results chan<- int,
) {
	for i := lowerBound; i <= upperBound; i++ {
		if num.IsPrime(i) {
			num.ProcessingPrimeNumbers(i)
			//fmt.Println("Robot", robotNumber, "found prime number:", i)

			results <- i
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()

	resultsChannel := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(6)

	go FindPrimeNumbers(1, 100, &wg, resultsChannel)
	go FindPrimeNumbers(101, 200, &wg, resultsChannel)
	go FindPrimeNumbers(201, 300, &wg, resultsChannel)
	go FindPrimeNumbers(301, 400, &wg, resultsChannel)
	go FindPrimeNumbers(401, 500, &wg, resultsChannel)
	go FindPrimeNumbers(501, 600, &wg, resultsChannel)

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	count := 0
	for primeNumber := range resultsChannel {
		count++
		fmt.Println("count increased from", count-1, "to", count, "with prime:", primeNumber)
	}

	fmt.Println("Total prime numbers found:", count)
	fmt.Println("finished in:", time.Since(start))
}

/*

- SYNCHRONIZATION
  - memory sharing => mutex
  - message passing => channel

- Which one to use?
  - Go’s philosophy on concurrency can be summed up like this:
    aim for simplicity, use channels when possible, and treat goroutines like a free resource.

*/
