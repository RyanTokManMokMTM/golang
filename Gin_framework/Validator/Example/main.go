package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

//reference : https://github.com/go-playground/validator
//All usable tag is marked there

type User struct {
	FirstName string `validate:"required,alphabetOnly"`
	LastName string `validate:"required"`
	Age uint8 `validate:"gte=0,lte=130""`
	Email string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"`
	Address []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City string `validate:"required"`
	Planet string `validate:"required"`
	Phone string `validate:"required"`
}

//set custom validator with custom name(tag)
//like gin => validator tag name call binding
var validate *validator.Validate


func main(){
	validate = validator.New()
	validate.RegisterValidation("alphabetOnly",validateAlphabet)
	validateStruct()
	validateVariable()
}

//use own validation function
func validateAlphabet(fl validator.FieldLevel) bool{
	reg := regexp.MustCompile(`[\w]+`)
	matched := reg.FindAllString(fl.Field().String(),-1) //check current field value
	fmt.Println(matched)
	if len(matched) != 1{
		return false
	}

	return true
}

func validateStruct(){
	address := &Address{
		City: "HongKong",
		Street: "New Territories",
		Planet: "Tsuen Wan",
		Phone: "none",
	}

	user := &User{
		FirstName: "jackson",
		LastName: "tmm",
		Age : 23,
		Email: "Jackson@hotmail.com",
		FavouriteColor: "#333",
		Address: []*Address{address},
	}

	//validate the struct with our validate tag
	if err := validate.Struct(user);err!=nil{
		//case err to validate error type
		//ok = true => cast succeed
		if _,ok := err.(*validator.InvalidValidationError);ok{
			fmt.Println(err)
			return
		}

		//if the struct just validation error ,print out all its property
		for _,err := range err.(validator.ValidationErrors){
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println(err.Kind())
		}
		return
	}
	fmt.Println("Struct validated successfully!Saving to db...")
	//validated ,save the data to somewhere
}

func validateVariable(){
	email := "jackson@hotmail.com"
	err := validate.Var(email,"required,email") //just validate current tag
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Email is valid")
}