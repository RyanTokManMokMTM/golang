package main

import (
	"fmt"
	dler "get-api/downloader"
	myfmt "get-api/fmt"
)

func main(){
	fmt.Println("hello world")
	allReq := dler.Request{[]string{"BV1Ff4y187q9", "BV13q4y1U7JU","BV1gq4y1D7R1"}}
	req,err := dler.Downloader(allReq); if err != nil{
		panic(err)
	}

	for _,dataInfo := range req.Result{
		myfmt.Logger.Printf("title %s\n desc :%s\n",dataInfo.Data.Title,dataInfo.Data.Desc)
	}
}
