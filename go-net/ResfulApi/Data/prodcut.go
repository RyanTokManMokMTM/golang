package Data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"'`
	Description string `json:"description"`
	Price float32 `json:"price"`
	CreatedOn string `json:"_"` //ignore by json decoder
	UpdatedOn string `json:"_"` //ignore by json decoder
	DeleteOn string `json:"_"` //ignore by json decoder
}

func (p *Product)FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

type Products []*Product

var ErrorNotFound = errors.New("not found")

func (p *Products)ToJSON(w io.Writer) error{
	//here we want to define our custom JSON Encode and store to w
	encoder := json.NewEncoder(w)
	return encoder.Encode(p) //it will ,write the json encoding with v
}

func UpdateProduct(id int,p *Product) error{
	_ , index , err := findProduct(id)
	if err != nil{
		return err
	}

	//if the product is found with the id
	//replace with return id
	p.ID = id
	productsList[index] = p
	return nil
}

func findProduct(id int) (*Product,int,error){
	for i, p := range productsList{
		if p.ID == id{
			return p,i,nil
		}
	}
	return nil,-1,ErrorNotFound
}

func GetProducts() Products{
	return productsList //get all products form db
}

func AddProducts(p *Product){
	p.ID = generator()
	productsList = append(productsList,p)
}

func generator() int {
	//here we get the id of the last element is the list +1
	last := productsList[len(productsList)-1]
	return last.ID + 1
}

var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}