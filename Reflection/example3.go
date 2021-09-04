//is reflection value can be set??
//reflect.canSet(),reflect.canAddr() only for pointer slice struct
package main

//
//func main(){
//	//這個🌰就是只有canSet以及canAddr 的才能改變值
//	//addr 的type dereferencing pointer struct slice
//	floatVar := 3.14
//	val := reflect.ValueOf(floatVar)
//	fmt.Printf("this can set? %v and this can addr? %v\n",val.CanSet(),val.CanAddr())
//	//val.SetFloat(6666.5565) // can't be set ,because not  addressable
//
//	floatVarAddr := 3.14
//	valF := reflect.ValueOf(&floatVarAddr)
//	valF = valF.Elem() //get the element to the r.value
//	fmt.Printf("this can set? %v and this can addr? %v\n",valF.CanSet(),valF.CanAddr())
//	valF.SetFloat(66.66666)
//	println(floatVarAddr)
//}