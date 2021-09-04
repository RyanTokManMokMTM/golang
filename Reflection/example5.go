//Use reflection to make other function to call
package main

import (
	"fmt"
	"reflect"
	"time"
)

func delay(){
	//sleep a second
	time.Sleep(1*time.Second)
}

func makeFunction(obj interface{}) interface{}  {
	//get obj value and type
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj) //if it is a function, it will be the function addr

	//check the type is reflection.Func?
	if objType.Kind() != reflect.Func{
		panic("Must be a function")
	}

	//new a new reflection func base on the type
	newFunction := reflect.MakeFunc(objType,func(args []reflect.Value) (result []reflect.Value){
		//this function to calculate the time
		start := time.Now()
		objValue.Call(args) //call our function
		end := time.Now()
		fmt.Printf("time is used %v",end.Sub(start))
		return
	})
	return newFunction.Interface()
}

//func main(){
//	timerFunc := makeFunction(delay).(func()) // change interface to func
//	timerFunc() //call the func
//}
