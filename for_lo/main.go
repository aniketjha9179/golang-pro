package main

import (
	"fmt"

)

func main() {
	for  i := 0; i <1; i++ {
		fmt.Println("Numbers is :", i)

	}
	// to run infinite loop like while loop
	counter:=0
	for{
		
		println("table of two is")
		fmt.Print("2 ka table",counter )
		
		counter++
		if counter==1{
			break
		}
	}
	number:=[]int{44,42,44,45,46,64,744}
	for index,value:=range number{
		fmt.Printf("Index of number is %d,and value is %d \n",index, value)
	}
	var data= "kittiputti"
	for index, value:=range data{
		fmt.Printf("value of index is %d and char is %c \n",index,value)
	}

}