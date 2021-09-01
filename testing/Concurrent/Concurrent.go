//BenchMarking Testing
//It is used to test the

package main

import (
	"sync"
	"sync/atomic"
)

func atomicCounter(counter *int64,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0; i<10000; i++{
		atomic.AddInt64(counter,1)
	}
}

func mutexCounter(counter *int64,wg *sync.WaitGroup,mux *sync.Mutex){
	defer wg.Done()
	for i:=0;i<10000;i++{
		mux.Lock()
		*counter++
		mux.Unlock()
	}
}

func ConcurrentAtomic() int64{
	wg := sync.WaitGroup{}
	wg.Add(2)
	var counter int64 = 0
	go atomicCounter(&counter,&wg)
	go atomicCounter(&counter,&wg)
	wg.Wait()
	return counter
}

func ConcurrentMutex() int64{
	wg := sync.WaitGroup{}
	wg.Add(2)
	var mux sync.Mutex
	var counter int64 = 0
	go mutexCounter(&counter,&wg,&mux)
	go mutexCounter(&counter,&wg,&mux)
	wg.Wait()
	return counter
}

func main(){
	println(ConcurrentAtomic())
	println(ConcurrentMutex())
}