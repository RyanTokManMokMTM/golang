package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request){
		fmt.Fprintf(rw,"%s %s %s \n",r.Method,r.URL,r.Proto) //send to response 
		for k,v := range r.Header{
			fmt.Fprintf(rw,"Header[%q] = %q\n",k,v) //send to response 
		}
		fmt.Fprintf(rw,"Host = %q\n",r.Host)
		fmt.Fprintf(rw,"remote addr :%q\n",r.RemoteAddr) //send to response 
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k,v := range r.Form{
			fmt.Fprintf(rw,"Form[%q] = %q \n",k,v)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}