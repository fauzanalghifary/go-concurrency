package main

import "fmt"

func double(numbers []int, doubleChannel chan int) {
	for _, number := range numbers {
		doubleChannel <- number * 2
	}

	close(doubleChannel)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubleChannel := make(chan int)

	go double(numbers, doubleChannel)

	for doubledNumber := range doubleChannel {
		fmt.Println("received", doubledNumber)
	}
}

/*

- range over a channel: loop until the channel is closed

*/
