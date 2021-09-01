package main

import (
	"fmt"
	"time"
	"sync"
)

func saySth(str string, waitGp *sync.WaitGroup) {
	defer waitGp.Done() // before end/after the loop run this code.Done() = waitGroup counter - 1
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

//Using WaitGroup
//wait the Group before counter = 0

//pros: no wasting time
//Cons: manually set  the counter.waitGp.add(number)

func main() {
	waitGp := new(sync.WaitGroup)
	waitGp.Add(2) // wait 2 routin
	//HERE is main goroutine
	//When this main goroutin is ended,all other will be killed/closed
	//Asynchronous
	go saySth("Hello",waitGp) //print 5 time first //run on other thread
	go saySth("World",waitGp) //print 5 time later //run on other thread

	(*waitGp).Wait() // wait until counter = 0

}
