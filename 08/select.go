package main

import (
	"fmt"
	"time"
)

func writeEvery(message string, seconds time.Duration, channel chan string) {
	for {
		time.Sleep(seconds * time.Second)
		channel <- message
	}
}

func main() {
	tickChannel := make(chan string)
	boomChannel := make(chan string)

	go writeEvery("tick", 2, tickChannel)
	go writeEvery("BOOM", 5, boomChannel)

	for {
		select {
		case msg1 := <-tickChannel:
			fmt.Println(msg1)
		case msg2 := <-boomChannel:
			fmt.Println(msg2)
		}
	}
}

/*

-

*/
