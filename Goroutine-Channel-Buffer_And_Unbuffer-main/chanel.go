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

/*
Channel 就想一條管子，goroutine 可以從管子中拉資料下來或者推資料到管子
Channel 的 block
1.在push資料到channel 會等待 有人pull這筆資料後才會繼續下一步
2.在pull的時候 如果channel 是空的/沒有資料可以等待就會等待,直到有goroutine push 才會進行

**Unbuffered Channel(only 1 value in a channel)
 Cons：只要有人要pull/push 都必須等待,如果push的執行很短而 pull的執行很長 ,push的人就必須等待到pull完才能繼續

** Buffered Channel(like a array collection)
	只有length被填滿才會被block住
	例如size = 5，目前channel 數值為5，有人要push 該push的goroutine 就會被block 住
	

*/
