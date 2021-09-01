package main

import (
	"log"
)

func main(){
	ch := make(chan int,10) //buffer size is 10

	go func(){
		for i:=0;i<10;i++{
			ch <- i
		}
		close(ch) //now the channel is no longer available
	}()

	for {
		value,isClose := <-ch
		if !isClose{ //when channel is closed break the loop
			break
		}

		log.Println(value)
	}

	//can not push value when channel is closed(panic)
	//ch <- 1 

	// time.Sleep(2*time.Second)

	// for i:= range ch{
	// 	log.Println(i)
	// }
	
}