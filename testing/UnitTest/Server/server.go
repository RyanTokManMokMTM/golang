package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter,r *http.Request){
	value := r.FormValue("val")
	if value == ""{
		//send error message
		http.Error(w,"required a value",http.StatusBadRequest)
		return
	}

	val ,err := strconv.Atoi(value) //convert string to int
	if err != nil{
		//not an int
		//send error message
		http.Error(w,"the value is not a int",http.StatusBadRequest)
		return
	}

	//send res to client
	//send error message
	if _,err := fmt.Fprintln(w,val*2);err != nil{
		http.Error(w,"can't write to response",http.StatusBadRequest)
		return
	}
}

func main(){
	//create a http server
	http.HandleFunc("/cal",Handler)
	err := http.ListenAndServe(":8080",nil);if err != nil{
		log.Fatalln(err)
	}

}
