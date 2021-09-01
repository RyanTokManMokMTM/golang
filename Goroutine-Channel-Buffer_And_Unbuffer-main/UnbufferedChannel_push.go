package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine Started.")
		time.Sleep(time.Second)
		fmt.Println("Goroutne Ended")

		ch <- "Go finished"
		fmt.Println("Channel value pulled")
	}()

	time.Sleep(time.Second * 2) //wait for 2 second
	fmt.Println(<-ch)
	fmt.Println("main goroutine is finished")
	//get data from channel
}

/*
Working logic:
main goroutine will waiting for 2s
goroutine 1(suppose is goroutine 1) will print
	started -> 1s later -> Ended -> push to channel ->waiting

2s later:
	main routine get the data from channel, at the mounent, goroutine 1 keep running and print the last string
	then main goroutine print the its own last string

*/

/*
//Output
Goroutine Started.
Goroutne Ended
Go finished
main goroutine is finished
Channel value pulled
*/
