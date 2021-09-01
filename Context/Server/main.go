//a simple server with http package
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request){
	/*
	func (r *Request) Context() context.Context {
		if r.ctx != nil {
			return r.ctx
		}
		return context.Background()
	}
	*/

	fmt.Println("Incoming request")
	ctx := r.Context() //return the root context or request ctx
	select {
	case <-time.After(5*time.Second):
		fmt.Println("the words is done")
	case <-ctx.Done():
		if err := ctx.Err();err!= nil{
			fmt.Println(err)
		}
	}
	fmt.Fprint(w,"hello,world")
	fmt.Println("Response request")
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}