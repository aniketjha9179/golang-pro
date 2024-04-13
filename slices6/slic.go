package main

import (
	"fmt"

)

func main() {
	
	// number:=[]int {1,2,4,4,5,6}
	// number=append(number,4,5,55,664,55444,3433 )
	// fmt.Println("numbers", number)
	// fmt.Printf("Number has data type of : %T", number)
	// fmt.Println(" Length:", len(number))
	//  age:=[]int{}
	//  fmt.Println(age)
	number:=make([]int, 3,5)
	number=append(number, 4)
	number=append(number, 88)
	number=append(number, 88)
	fmt.Println("Slice",number)
	fmt.Println("length",len(number))
	fmt.Println("capacity",cap(number))
	

}
