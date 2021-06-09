package main

import (
	"fmt"
	// "os"
	// "strconv"
) 

func main() {
	fmt.Println("Mutable & Immutable Data Types")
	//Mutable & Immutable Data Types
	// x:=5
	// y:=x
	// y=7
	// fmt.Println(x,y)
	// x:=[]int{3,5,6}
	// y:=x
	// y[0]=100
	// fmt.Println(x,y)
	//MAPS
	// x:=map[string]int{"hello":3}
    // y:=x
	// y["y"]=100
	// x["x"]=7
	// fmt.Println(x,y)
	//Array
	x:=[2]int{4,6}
	y:=x
	y[0]=8
	fmt.Println(x,y)
}