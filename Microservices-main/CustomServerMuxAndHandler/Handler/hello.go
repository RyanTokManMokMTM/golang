package Handlers

import (
	"fmt"
	"log"
	"net/http"
)

//need to implement handler interface

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

type Hello struct {
	//here define a logger for this hello to use
	logger *log.Logger
}

func NewHelloHandler(log * log.Logger) *Hello{
	return &Hello{
		logger: log,
	}
}

func (ser *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	log.Println("Hello Custom Handler")
	fmt.Fprintln(rw,"Hello")
}