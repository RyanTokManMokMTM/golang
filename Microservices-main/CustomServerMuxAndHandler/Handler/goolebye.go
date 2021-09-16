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

type GoodBye struct {
	//here define a logger for this hello to use
	logger *log.Logger
}

func NewGoodByteHandler(log * log.Logger) *GoodBye{
	return &GoodBye{
		logger: log,
	}
}

func (server *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	log.Println("GoodByte Custom Handler")
	fmt.Fprintln(rw,"GoodBye")
}