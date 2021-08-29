package main
import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"URL Path = %q\n",r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
