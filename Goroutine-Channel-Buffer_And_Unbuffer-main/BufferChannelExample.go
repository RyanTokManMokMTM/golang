package main
import (
	"log"
)

func main(){
	ch := make(chan int,1)// length is 1
	ch <- 1 //add value 1 to channel and current size is not greater than channel size won't block
	log.Println(<-ch)
}
