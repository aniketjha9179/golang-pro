package main

import "fmt"

func modifyValueByRefrence(num *int){
	*num =*num +5
}

func main() {
	num := "AN"
	ptr := &num
	// fmt.Println("Num has value", num)
	fmt.Println("ptr memory location", ptr)
	fmt.Println("Data contains thorugh pointer", *ptr )
	value:=10

	modifyValueByRefrence(&value)
	fmt.Println("value contains", value)
	


}