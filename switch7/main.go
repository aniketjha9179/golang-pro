package main

import "fmt"

func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("tuedsay")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thrusday")
	case 5:
		fmt.Println("friday")
	default:
		fmt.Println("saturday")
	}
	month:="january"

	switch month{

	case "january", "February", "March":
		fmt.Println("Winter")
	case "April", "may", "june":
		fmt.Println("summer")
	default :
		fmt.Println("Other Seasons")
	}
	


}
