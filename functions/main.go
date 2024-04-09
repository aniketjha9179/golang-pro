package main

import "fmt"


func printFunction(){
	fmt.Println("THis functioon will we called inside main func")
}

func multiPleOftwo(a , b int)(ans int ){
	return a*b

}
func main(){
	println("Calling first function" )
	printFunction()
	ans:=multiPleOftwo(3,3)
	fmt.Println("Multiple of two number is:" , ans)
}