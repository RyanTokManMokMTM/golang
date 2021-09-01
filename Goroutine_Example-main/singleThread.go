package main

import (
	"fmt"
	"time"
	"runtime"
)

func saySth(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(str)
	}
}

func main() {
	//Synchronous
	saySth("Hello") //print 5 time first
	saySth("World") //print 5 time later
}


// func main() {
// 	//HERE is main goroutine
// 	//When this main goroutin is ended,all other will be killed/closed
// 	//Asynchronous
// 	go saySth("Hello") //print 5 time first //run on other thread
// 	go saySth("World") //print 5 time later //run on other thread
// 	fmt.Print(runtime.NumCPU()) //run on main //run on other thread
	
// }
