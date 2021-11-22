package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	var isShow bool
	flag.BoolVar(&isShow,"version",false,"Version information")
	flag.BoolVar(&isShow ,"v",false,"Version information")
	flag.Parsed()
	if isShow {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	}
}
