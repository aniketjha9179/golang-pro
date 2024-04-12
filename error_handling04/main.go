package main

import (
	"fmt"
)


func divide(a,b float64)(float64, error){
	if b==0 {
		return 0,fmt.Errorf("denominator must not be zero")

	}
	return 	a/b, nil

	

}


func main(){
	fmt.Println("Started Error handling in the functions")
	ans, _:=divide(10,3)
	fmt.Println("Division of two numbers is " , ans)

}