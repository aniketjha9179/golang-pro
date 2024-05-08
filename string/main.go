package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Hellobro. bro . kaisa . hai .aniket"
	parts := strings.Split(data,".")
	fmt.Println(parts)
	str:="one two two two three four five six seven"
	count:=strings.Count(str, "tt")
	fmt.Println("Count is", count)
	str="                  hello ,                    go !   "
	trimmed:=strings.TrimSpace(str)
	fmt.Println("Trimmed", trimmed)

	str1:="Aniket"
	str2:="Jha"
	result:=strings.Join([]string{str1, str2}, " ")
	fmt.Println("Results", result)

}
