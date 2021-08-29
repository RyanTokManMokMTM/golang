//using for to print arg
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//using for to read the Args Slice
	for _, arg := range os.Args[1:] { //return index and element ,using the range
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
