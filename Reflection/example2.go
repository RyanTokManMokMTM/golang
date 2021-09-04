package main

type Student struct {
	Name string
}

//func main(){
//	//Convert reflection type a type
//	fVar := 3.14
//	val := reflect.ValueOf(fVar)
//
//	//change value to interface
//	newFVal := val.Interface().(float64) //change interface to float64 interface
//	fmt.Println(newFVal + 1.1)
//	//append to a slice
//	sliceVar := make([]int,5)
//	val = reflect.ValueOf(sliceVar) //reflect our slice
//
//	//or can append reflect value with reflect value
//	val = reflect.Append(val,reflect.ValueOf(6))
//	newSlice := val.Interface().([]int) //change reflect value to []int slice
//	newSlice = append(newSlice, 5) //append to slice
//	fmt.Println(newSlice)
//
//	//convert reflect value to a struct
//	//new a reflection with a type
//	myStu := reflect.New(reflect.TypeOf(Student{})) //a pointer /return a reflect value with reflect type
//	//get ele
//	stu := myStu.Elem() //get the value of that pointer
//	nameField := stu.FieldByName("Name")
//	if nameField.IsValid(){
//		//check it field can set or not
//		//canSet return canAddr() ,only addressable can set
//		if nameField.CanSet(){
//			nameField.SetString("Jackson")
//		}
//		//change reflection to our struct
//		result := stu.Interface().(Student)
//		fmt.Println(result)
//	}
//}