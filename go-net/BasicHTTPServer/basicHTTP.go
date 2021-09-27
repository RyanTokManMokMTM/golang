package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	//using default server mux
	//HandleFunc will call default server mux handler set to a map
	http.HandleFunc("/hello",func(rw http.ResponseWriter,r*http.Request){
		log.Println("Hello ,welcome")
		fmt.Fprint(rw,"welcome") //write the message to response
	})

	http.HandleFunc("/",func(rw http.ResponseWriter,r *http.Request){
		log.Println("Welcome to basic http handler")

		//suppose is a post function
		bytes , err := io.ReadAll(r.Body)
		if err != nil{
			log.Println("Error reading",err)
			http.Error(rw,"unable to read body",http.StatusBadRequest)
			return
		}
		fmt.Printf("Hello %s",bytes)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("goodbye")
		fmt.Fprintln(rw,"good byte")
	})

	log.Println("Starting server...")
	log.Fatalln(http.ListenAndServe(":8080",nil)) //using default server mux
}