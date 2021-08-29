package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) //read the input
	for input.Scan() {                  //scan each line
		counts[input.Text()]++
	}

	for line, txt := range counts {
		if txt > 1 {
			fmt.Printf("%d\t%s\n", txt, line)
		}
	}
}
