//simple client
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main(){
	ctx := context.Background()
	ctx,cancel := context.WithTimeout(ctx,time.Duration(8*time.Second)) //5 seacond
	defer cancel() //if Timeout cancel()

	//create a new request
	req, err:= http.NewRequest("get","http://localhost:8080/",nil)
	if err != nil{
		log.Fatal(err) //os.exist
	}

	//put the content to req
	req = req.WithContext(ctx)

	//send the request
	res,err := http.DefaultClient.Do(req) //current client send req
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	resBytes,err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%s",resBytes)
}