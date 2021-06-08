package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
func test(x,y int)(z1 int,z2 int){
	defer fmt.Println("clean up the extra memory use")
	z1=x+y
	z2=y-x
	fmt.Println("start")
	return 
}
func main() {
	fmt.Println("Functions")
	//Functions
	add,sub:=test(4,9)
	fmt.Println("Sum:",add,sub)
}