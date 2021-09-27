package APIHandler

import (
	"fmt"
	"go-net/ResfulApi/Data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProduct(logs *log.Logger) *Products{
	return &Products{logger: logs}
}

//implement the Handler
func (p *Products) ServeHTTP(rw http.ResponseWriter,r *http.Request){

	//check is get method
	if r.Method == http.MethodGet{
		p.getProducts(rw,r)
		return
	}else if r.Method == http.MethodPost{
		p.logger.Println("Handel POST REQUEST")
		p.addProducts(rw,r)
		return
	}else if r.Method == http.MethodPut{
		p.logger.Printf("Handel PUT REQUEST %s",r.URL.Path)
		uri := r.URL.Path

		//get the id from uri with regex
		//match with regex one or more 0-9 number
		reg := regexp.MustCompile(`/([0-9]+)`)
		resultStr := reg.FindAllStringSubmatch(uri,-1)

		fmt.Println(resultStr)
		if len(resultStr) != 1{
			//not exactly 1 result(we only get 1 id)
			p.logger.Println("Invalid URI more than 1 id")
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}

		if len(resultStr[0]) != 2 {
			//result string[][] at least contain 2 element
			p.logger.Println("Invalid URI more then one capture group")
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}

		//get the string
		idStr := resultStr[0][1]

		//convert str to int
		id , err := strconv.Atoi(idStr)
		if err != nil{
			p.logger.Println("Invalid URI: Unable to convert to num")
			http.Error(rw,"Invalid URI",http.StatusBadRequest)
			return
		}

		p.updateProducts(id,rw,r)
		return
	}
	//post and put

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

//Feature about the Product
func (p *Products) getProducts(rw http.ResponseWriter,r *http.Request){
	log.Println("Welcome! GET METHOD AND GIVE U ALL THE PRODUCTS WE HAVE")

	//get the product from our data
	productsList := Data.GetProducts()

	//using the custom encoder ,marshal our data to json with our interface with json tag
	if err := productsList.ToJSON(rw);err != nil{
		http.Error(rw,"Unable to marshal json",http.StatusBadRequest)
	}
}

//add the product
func (p *Products) addProducts(rw http.ResponseWriter,r *http.Request){
	//we need to generate the product id
	data := &Data.Product{}

	if err:= data.FromJSON(r.Body);err != nil{
		http.Error(rw,"Unable to unmarshal json",http.StatusBadRequest)
	}

	//add product to the db list
	Data.AddProducts(data)
}

func (p *Products) updateProducts(id int ,rw http.ResponseWriter,r *http.Request){
	//get data from body
	data := &Data.Product{}
	if err := data.FromJSON(r.Body);err != nil{
		p.logger.Println("Unable to unmarshal json")
		http.Error(rw,"Unable to unmarshal json",http.StatusBadRequest)
		return
	}

	fmt.Println(data)

	//update the list
	err := Data.UpdateProduct(id,data)
	if err == Data.ErrorNotFound{
		http.Error(rw,"Product is not found",http.StatusNotFound)
		return
	}

	if err != nil{
		http.Error(rw,"Unable to update the product!",http.StatusInternalServerError)
		return
	}

}