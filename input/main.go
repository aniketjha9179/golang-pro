package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	fmt.Println("Enter your name ")
	var name string
	// fmt.Scan((&name))
	// how to get  long string as input so we use bufio
	reader:=bufio.NewReader(os.Stdin)
	name,_=reader.ReadString('\n')
	fmt.Println("Hello Mr", name)

	


}
