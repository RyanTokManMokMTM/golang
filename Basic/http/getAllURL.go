package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string,ch chan <- string){
	start := time.Now()
	res,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//if get url
	//copy to ioutil.Discard will always success,without any reason
	//just return the bytes,
	nbytes , err := io.Copy(ioutil.Discard,res.Body)
	res.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs %7d %s",secs,nbytes,url)
	
}

func main(){
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		fmt.Println(<- ch) //will stuck here and wait for channal input from goroutine
	}

	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}