package main
import(
	"log"
)

func main(){
	ch := make(chan int)
	ch <- 1 //always waiting here(main goroutine asleep)->deadlock
	log.Println(<- ch)
}