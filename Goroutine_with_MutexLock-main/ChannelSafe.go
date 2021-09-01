package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0
	ch := make(chan int, 1)
	ch <- total

	for i := 0; i < 1000; i++ {
		go func() {
			// pull fron ch
			ch <- <-ch + 1 // = <-ch + 1 //get channel value and plus 1,then ch <- =>push new value to channel
		}()
	}
	time.Sleep(time.Second) //wait for a second
	fmt.Println(<-ch)       //get final value from
}

//Channel data if pull by other goroutine ,will make another gorouitine stop and keep waiting
//after the Channel has value and the other will continues its calculation
//if it is always keep waiting,-> may having deadlock!!
