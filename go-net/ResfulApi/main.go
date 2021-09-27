package main

import (
	"context"
	"fmt"
	Handlers "go-net/ResfulApi/APIHandler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	logs := log.New(os.Stdout,"restfulApi:",log.LstdFlags)

	apiHandler := Handlers.NewProduct(logs)
	serMux := http.NewServeMux()

	//just handle 1 route
	serMux.Handle("/",apiHandler)

	server := http.Server{
		Addr: ":8080",
		Handler: serMux,
		ErrorLog: logs,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
		ReadHeaderTimeout: 5*time.Second,
		IdleTimeout: 5*time.Second,
	}

	go func() {
		fmt.Println("Server is listening")
		log.Fatalln(server.ListenAndServe())
	}()

	ch := make(chan os.Signal)
	defer close(ch)

	signal.Notify(ch,os.Interrupt)
	signal.Notify(ch,os.Kill)
	osSignal := <- ch
	log.Println(osSignal)

	ctx ,_ := context.WithTimeout(context.Background(),30 * time.Second)
	server.Shutdown(ctx)
}
