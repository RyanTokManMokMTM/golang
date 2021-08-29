package fmt

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init(){
	//open a file
	file,err := os.OpenFile("log.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,666)
	if err != nil{
		panic(err)
	}
	Logger = log.New(io.MultiWriter(os.Stdout,file),"log: ",log.LstdFlags)
}
