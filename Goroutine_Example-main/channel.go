package main

import (
	"fmt"
	"time"
)

func saySth(str string, myChan chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(str)
	}
	myChan <- 0
}

/*
Before main goroutine is ended,need to wait other routine finished
->using sleep
->using waitGroup
->using channel
*/

//Using Channel(main feature in go)

//pros: no wasting time
//Cons: manually set  the counter.waitGp.add(number)

func main() {
	myChan := make(chan int)

	//HERE is main goroutine
	//When this main goroutin is ended,all other will be killed/closed
	//Asynchronous
	go saySth("Hello", myChan) //print 5 time first //run on other thread
	go saySth("World", myChan) //print 5 time later //run on other thread

	id := <-myChan //waiting to get data from channel =>otherwise ->deallock
	fmt.Printf("first %d \n", id)

	id_2 := <-myChan //waiting to get data from channel =>otherwise ->deallock
	fmt.Printf("second %d \n", id_2)

}
