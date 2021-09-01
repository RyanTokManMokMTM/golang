package main

import (
	"fmt"
	"sync"
)
var group sync.WaitGroup
func main(){
	group.Add(2)
	go func(str string){
		fmt.Println(str)
		group.Done()
	}("123")
	go func(str string){
		fmt.Println(str)
		group.Done()
	}("abc")
	group.Wait()
}
