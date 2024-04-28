package main

import "fmt"

func main() {
	studentsGrades := make(map[string]int)
	studentsGrades["Prince"] = 33
	studentsGrades["rachit"] = 77
	studentsGrades["Bob"] = 55
	studentsGrades["sam"] = 99
	fmt.Println("Marks of sam", studentsGrades["sam"])
	studentsGrades["Bob"]=100
	// delete(studentsGrades,"sam");
	fmt.Println("Marks of Bob", studentsGrades["sam"])
	// checking if a key is exists 

	grades , exists :=studentsGrades["Prince"]
	fmt.Println("Grades of Prince :", grades)
	fmt.Println("Prince exists:",exists)
fmt.Println("makrs of Prince",studentsGrades["Prince"])

for index, value:=range studentsGrades{
	fmt.Printf("Key is  %s and marks is %d\n",index, value)
}

person:=map[string]int{
	"alice":45,
	"kut":53,
	"nick":44,
}
for index,value:=range person{

	fmt.Printf("Key is %s, and marks is %d\n", index, value)
}




}