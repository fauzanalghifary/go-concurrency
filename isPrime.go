package go_concurrency

import "time"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func ProcessingPrimeNumbers(n int) {
	// Simulate a long-running operation
	time.Sleep(100*time.Millisecond + time.Duration(n)*0*time.Millisecond)
}
