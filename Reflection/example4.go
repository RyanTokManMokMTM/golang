//reflection method
package main

import (
	"fmt"
)

type Friend struct {
	Name string
}

func (f Friend) PlayGame(gameA string,gameB string){
	fmt.Printf("i am playing with %v and playing with %s and %s",f.Name,gameA,gameB)
}

//func main(){
//	fd := reflect.New(reflect.TypeOf(Friend{}))
//	fdV := fd.Elem()
//	fdName := fdV.FieldByName("Name")
//	if fdName.IsValid(){
//		if fdName.CanSet(){
//			fdName.SetString("jackson")
//		}
//	}
//
//	//can also can call its func
//	fdMethod := fdV.MethodByName("PlayGame")
//	if fdMethod.IsValid(){
//		//the method has no argument
//		fdMethod.Call([]reflect.Value{reflect.ValueOf("Gta"),reflect.ValueOf("AVA")})
//	}
//}