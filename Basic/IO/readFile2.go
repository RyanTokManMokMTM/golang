package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	for _, name := range file {
		data, err := ioutil.ReadFile(name) //data is bytes[]
		if err != nil {
			fmt.Fprintf(os.Stderr, "readFileds %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			//we can get data(name of the file and \n)
			//each line
			fmt.Println(line)
			counts[line]++
		}
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
