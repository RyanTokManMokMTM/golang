package main

import (
	"context"
	"go-net/CustomServerMuxAndHandler/Handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	//create custom servermux
	logs := log.New(os.Stdout,"handler :",log.LstdFlags)
	a := Handlers.NewHelloHandler(logs)
	b := Handlers.NewGoodByteHandler(logs)
	serMux := http.NewServeMux()

	//set the handler to server mux
	serMux.Handle("/",a)
	serMux.Handle("/goodbye",b)

	//set up the serverMUX
	server := http.Server{
		Addr: ":8080", //configure the bind address
		Handler: serMux, // set the handler
		ErrorLog: logs, //set the error logger
		ReadTimeout: 5 * time.Second, //time to read the request
		WriteTimeout: 10 * time.Second, //time limit to write to the client
		IdleTimeout: 120 * time.Second, //time to keep connection alive
	}

	go func(){
		log.Println("server is starting...")
		log.Fatalln(server.ListenAndServe())
	}()

	//using a goroutine to listing the server
	//so using a channel to receive a signal from os
	ch := make(chan os.Signal)
	signal.Notify(ch,os.Interrupt)
	signal.Notify(ch,os.Kill)

	//get the signal if it has sent to the channel by signal notify
	osSig := <- ch // stuck if channel have any signal yet
	log.Println(osSig)

	//if our server have any requests haven't finished
	//it will allow it to finish and close to server(not allowing any request anymore)
	//using context to close the request within the time
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second) //timeout after 30s
	server.Shutdown(ctx) //if timeout or the server pollTimeInterval is sent to the channel , it will close
}