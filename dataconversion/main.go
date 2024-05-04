package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 43
	fmt.Println("Number is ", num)
	// num=num+1.33
	fmt.Printf("type of num is %T\n", num)
	data:=float64(num)
	data=data+1.23
	fmt.Println("Data is", data)
	fmt.Printf("Data is %T\n ", data)
	num=123
	str:=strconv.Itoa(num)
	fmt.Println("TYPE IS ", str)
	fmt.Printf("Data is %T\n ", str)
	number:="12434"
	number_int,_:=strconv.Atoi(number)
	fmt.Printf("Type of number_int %T", number_int)
	// fmt.Println("TYPE IS ", number_int)

	num_string:="3.23"
	num_float,_:=strconv.ParseFloat(num_string, 64)
	fmt.Println("Type of number_int ", num_float)
	fmt.Printf("Type of number_int %T\n", num_float)






}