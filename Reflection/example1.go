////simple example usage about reflection
package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
////Reflection allow as to detect what type is running on
////and do different thing on that type
////such as struct to print struct fields,func call that etc
//func metaData(obj interface{}){
//	// type: the type of obj
//	// value : the value of current type
//	// kind : the primitive type(Prototype)
//	// Name : the name of current interface
//	ot := reflect.TypeOf(obj)
//	ov := reflect.ValueOf(obj)
//	oKind := ot.Kind()
//	oName := ot.Name()
//
//	fmt.Printf("current Object meta: type:%s,name:%s,kind:%s,value:%s\n",ot,oName,oKind,ov)
//}
//
//type MyFunc func(int,int) int
//func main(){
//	intVar := 10
//	strVar := "reflection"
//	metaData(intVar)
//	metaData(strVar)
//
//
//	type School struct{
//		name string
//		city string
//		year int
//	}
//
//	 mySchool := School{
//	 	name:"CGU",
//	 	city:"TW",
//	 	year:30,
//	 }
//	 metaData(mySchool)
//
//	 var total MyFunc = func(a,b int) int{
//	 	return a+b
//	 }
//
//	 metaData(total)
//
//	 testFunc := func(a,b int) int{
//	 	return a+b
//	 }
//
//	 metaData(testFunc)
//}
