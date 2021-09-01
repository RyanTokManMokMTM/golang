package main

import (
	"fmt"
	"time"
)

func saySth(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(str)
	}
}

/*
Before main goroutine is ended,need to wait other routine finished
->using sleep
->using waitGroup
->using channel
*/

//Using time
//Cons:we don't know how long need to sleep

func main() {
	//HERE is main goroutine
	//When this main goroutin is ended,all other will be killed/closed
	//Asynchronous
	go saySth("Hello") //print 5 time first //run on other thread
	go saySth("World") //print 5 time later //run on other thread

	time.Sleep(1000 * time.Millisecond) //sleep 1s and continue

}
