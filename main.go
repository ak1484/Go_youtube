package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
type Point struct{
	x int
	y int
	// z bool
}
type Circle struct{
	radius float64
	// center *Point
	*Point
}

func main() {
	fmt.Println("Structs and Custom Types")
	//Structs and Custom Types
	// var x Point=Point{1,2}
	// X:=Point{1,4}
	// fmt.Println(x)
	// fmt.Println(X)
	// x:=&Point{y:5}
	// y:=&x
	// fmt.Println(*y)
	// fmt.Println(x)
    // x.x=7
	// fmt.Println(*x)
	// p1:=&Point{2,4}
	// c1:=Circle{3.5,p1}
	c1:=Circle{3.5,&Point{2,9}}
	// fmt.Println(c1.Point)
	fmt.Println(c1.x)
}