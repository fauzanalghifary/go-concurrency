package main

import (
	"fmt"
)

func sum(numbers []int, sumChannel chan int) {
	total := 0
	for _, number := range numbers {
		total += number
	}

	sumChannel <- total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sumChannel := make(chan int)

	go sum(numbers, sumChannel)

	total := <-sumChannel
	fmt.Println("Total sum:", total)
}

/*

- CHANNELS
  - blocking: send and receive operations block until the other side is ready
  - might cause deadlock if not handled properly

*/
