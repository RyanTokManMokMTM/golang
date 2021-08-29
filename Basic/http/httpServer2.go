package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex
var count int

func main(){
	http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request){
		mux.Lock()
		count++ //this a  gorountine need to use mutex
		mux.Unlock()
		fmt.Fprintf(rw,"URL Path %q\n",r.URL.Path)
	})

	http.HandleFunc("/count",func(rw http.ResponseWriter, r*http.Request){
		mux.Lock()
		fmt.Fprintf(rw,"Counter  %d\n",count)
		mux.Unlock()
	})

	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}