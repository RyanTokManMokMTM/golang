package main
import (
	"log"
)

func main(){
	ch := make(chan int,10)
	go func(){
		for i:=0;i<10;i++{
			ch <- i*5
		}
		close(ch)
	}()

	for i:= range ch{ //use range to loop over channel and end at channel is closed
		log.Println(i)
	}
}