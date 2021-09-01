package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine is started waiting")
		time.Sleep(time.Second)
		fmt.Println("Goroutine is ended waiting")

		ch <- "Goroutine finished"
		fmt.Println("Goroutine pushed the value to channel")
	}()

	fmt.Println("Main goroutine is waiting for channel value")
	fmt.Println(<-ch) //pull from channel
	fmt.Println("Main goroutine received the value from channel")
}

/*
TODO:
Main routine print waiting message and sleep/wait
Goroutine 1 print started msg -> wait 1s -> ended msg -> push value to channel ->wait to pull

Main routine get the value and goroutine 1 continues to run -> print the channel value ->last msg in main routine
*/
