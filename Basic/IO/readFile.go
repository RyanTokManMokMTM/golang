package main
import(
	"bufio"
	"fmt"
	"os"
)

func countLine(f *os.File,count map[string]int){
	input := bufio.NewScanner(f) //os.Stdin is os.file interafce
	for input.Scan() {
		count[input.Text()]++
	}
}

func main(){
	counts := make(map[string]int)
	file := os.Args[1:] //file name from cmd
	if len(file) == 0 {
		countLine(os.Stdin,counts)
		//let use to input
	}else{
		//loop over the arg list
		for _, arg := range file{
			f,err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stdout,"readFile:%v\n",err)
				continue
			}
			countLine(f,counts)
			f.Close()
		} 
	}
	for line , value := range counts {
		if value > 1 {
			fmt.Printf("%d\t%s\n",value,line)
		}
	}
}