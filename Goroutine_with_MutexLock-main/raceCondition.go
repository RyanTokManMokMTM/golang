package main

import (
	"fmt"
	"time"
)

//Using goroutine/multiple thread

func main() {
	total := 0
	for i := 0; i <1000; i++ {
		go func() {
			total++ //may happen race condition
		}()
	}
	time.Sleep(time.Second) 
	fmt.Println(total)
}
