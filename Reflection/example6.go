//using reflection output struct
package main

import (
	"fmt"
	"reflect"
)

type People struct {
	name string
	age int
}

func (p People) PrintName(friend string,class int){
	fmt.Printf("This people name is %s\n",p.name)
}

func (p People) PrintAge(){
	fmt.Printf("This people name is %s\n",p.name)
}

type testFunc func(int,int) int

var (
	self = People{
		name: "jackson",
		age : 18,
	}

	calAdd testFunc = func(a,b int) int{
		return a + b
	}

	calM testFunc =  func(a,b int) int{
		return a * b
	}
)


func getMata(obj interface{}){
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	//here will check the reflection type(Kind)
	//Struct -> field and method : name in out
	//Func -> name in out
	if objType.Kind() == reflect.Struct{
		//get all field of the struct and func of struct
		fmt.Println("--------Print struct Fields------------")
		for id:=0;id<objType.NumMethod();id++{
			field := objType.Field(id) //get the field from reflection type
			value := objValue.Field(id) //get value from reflection value
			fmt.Printf("\t %v :\t %v\n",field.Name,value)
		}
		fmt.Println("--------Print struct Method------------")
		for id:=0;id<objType.NumMethod();id++{
			method := objType.Method(id) //get method from reflection type
			fmt.Printf("Struct method name %v and inputNum %v amd outputNum %v\n",
				method.Name,
				method.Type.NumIn(),
				method.Type.NumOut())
		}
	}else if objType.Kind() == reflect.Func{
		in := objType.NumIn() //get method from reflection type
		out := objType.NumOut() //get method from reflection type
		name := objType.Name() //get method from reflection type

		fmt.Printf("name %v and inputNum %v amd outputNum %v\n",
			name,
			in,
			out)
	}
}

func main(){
	getMata(self)
	getMata(calAdd)
	getMata(calM)
}
