package main
import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main(){
	url := os.Args[1]
	res,err := http.Get(url)
	if err != nil{
		fmt.Fprintf(os.Stderr,"fectch %v\n",err)
		os.Exit(1)
	}
	body , err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr,"fetch reading %s : %v \n",url,err)
		os.Exit(1)
	}
	fmt.Printf("%s",body)
}