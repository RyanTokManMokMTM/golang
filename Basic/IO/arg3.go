//Training arg
package main

import(
	"fmt"
	"os"
	"strings"
)

func main(){
	//out put the arg
	for i,arg := range os.Args {
		fmt.Printf("index %v is %v\n",i,arg)
	}
	fmt.Println(strings.Join(os.Args[1:],  " "))
}