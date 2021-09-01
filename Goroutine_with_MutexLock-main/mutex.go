package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	total int
	mux   sync.Mutex
}

func main() {
	total := Counter{total: 0}
	for i := 0; i <1000; i++ {
		go func() {
			total.mux.Lock()
			total.total++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.total)
	total.mux.Unlock()
}
