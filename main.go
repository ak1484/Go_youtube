package main

import (
	"fmt"
	// "os"
	// "strconv"
)
func main() {
	fmt.Println("If, Else If, Else ")
	//for if else condition
	age:=14
	if age>=18{
		fmt.Printf("You can drive")
	}else if age>=14{
		fmt.Printf("You can drive but with your parents")
	}else{
		println("you can not drive")
	}
}